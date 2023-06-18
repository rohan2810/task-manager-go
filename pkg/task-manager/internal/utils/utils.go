package utils

import (
	"encoding/json"
	"errors"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/types/known/timestamppb"
	v1 "rohan.com/task-manager/apis/v1/generated"
	"rohan.com/task-manager/pkg/task-manager/repo/models"
)

func taskMapper(dbTask *models.Task) (*v1.Task, error) {
	return &v1.Task{
		Id:               dbTask.ID,
		Name:             dbTask.Name,
		Description:      dbTask.Description,
		TaskCreator:      ConvertDBUserToProto(&dbTask.TaskCreator),
		Assignee:         ConvertDBUserToProto(&dbTask.TaskAssignee),
		TaskCreateTime:   timestamppb.New(dbTask.CreateTime),
		TaskDeadlineTime: timestamppb.New(dbTask.Deadline),
		Tags:             dbTask.Tags,
		Attachments:      ConvertDBAttachmentsToProto(dbTask.Attachments),
		//Subtasks:         taskMapperArray(dbTask.SubTasks),
		Completed:     dbTask.Completed,
		CompletedTime: timestamppb.New(dbTask.CompletedTime),
		Archived:      dbTask.Archived,
	}, nil
}

func ConvertDBUserToProto(dbUser *models.User) *v1.User {
	protoUser := &v1.User{
		Username: dbUser.Username,
		Fullname: dbUser.Username,
		Email:    dbUser.Email,
		Role:     dbUser.Role,
	}

	return protoUser
}

func ConvertDBAttachmentsToProto(attachments []*models.Attachment) []*v1.Attachment {
	attachmentsArray := make([]*v1.Attachment, len(attachments))
	for _, attachment := range attachments {
		attachmentsArray = append(attachmentsArray, &v1.Attachment{
			Name: attachment.Name,
			Type: attachment.Type,
			Uri:  attachment.URI,
		})
	}
	return attachmentsArray
}

func TaskMapperArray(dbTasks []*models.Task) []*v1.Task {
	tasksArray := make([]*v1.Task, len(dbTasks))
	for _, dbTask := range dbTasks {
		tasksArray = append(tasksArray, &v1.Task{
			Id:               dbTask.ID,
			Name:             dbTask.Name,
			Description:      dbTask.Description,
			TaskCreator:      ConvertDBUserToProto(&dbTask.TaskCreator),
			Assignee:         ConvertDBUserToProto(&dbTask.TaskAssignee),
			TaskCreateTime:   timestamppb.New(dbTask.CreateTime),
			TaskDeadlineTime: timestamppb.New(dbTask.Deadline),
			Tags:             dbTask.Tags,
			Attachments:      ConvertDBAttachmentsToProto(dbTask.Attachments),
			//Subtasks:         taskMapper(dbTask.SubTasks),
			Completed:     dbTask.Completed,
			CompletedTime: timestamppb.New(dbTask.CompletedTime),
			Archived:      dbTask.Archived,
		})
	}
	return tasksArray
}

func TransformMessage(src, dst interface{}) (interface{}, error) {
	var err error

	var dstM, ok = dst.(protoreflect.ProtoMessage)

	if !ok {
		return dst, errors.New("invalid conversion")
	}
	err = transformMessageToAPI(src, dstM)
	if err != nil {
		return dst, err
	}

	return dst, err
}

func transformMessageToAPI(src interface{}, dst protoreflect.ProtoMessage) error {

	srcBytes, err := json.Marshal(src)
	if err != nil {
		return err
	}
	marshaller := protojson.UnmarshalOptions{
		DiscardUnknown: true,
	}
	err = marshaller.Unmarshal(srcBytes, dst)

	if err != nil {
		return err
	}
	return nil
}
