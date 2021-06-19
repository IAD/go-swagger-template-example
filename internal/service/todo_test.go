package service

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/IAD/go-swagger-template-example/pkg/client/todoclient"
	"github.com/IAD/go-swagger-template-example/pkg/client/todoclient/todos"
	"github.com/IAD/go-swagger-template-example/pkg/client/todoclientmodels"
	"github.com/go-openapi/swag"
	fuzz "github.com/google/gofuzz"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/suite"
)

func TestTodo(t *testing.T) {
	t.Parallel()

	// nolint: exhaustivestruct
	suite.Run(t, &Suite{})
}

type Suite struct {
	suite.Suite

	client *todoclient.Todo
	logger *logrus.Entry
	fuzzer *fuzz.Fuzzer
}

func (s *Suite) SetupTest() {
	logger := logrus.NewEntry(logrus.New())

	logger.Info("starting")

	service, err := PrepareServer(0, logger)
	s.Require().NoError(err)
	s.Require().NotNil(service)

	go func() {
		errInt := service.Serve()
		if errInt != nil {
			logger.Error(errInt)
		}
	}()

	time.Sleep(time.Second)

	client := todoclient.NewClientWithBasePath(fmt.Sprintf("%s:%v", "localhost", service.Port), "")

	s.fuzzer = fuzz.New().NilChance(0).Funcs(func(i *string, c fuzz.Continue) {
		*i = fmt.Sprintf("%s*", c.RandString())
	})
	s.logger = logger
	s.client = client
}

func (s *Suite) TestAddOneHandler() {
	var item Item
	s.fuzzer.Fuzz(&item)

	var id int64
	s.fuzzer.Fuzz(&id)

	addOneCreated, addOneInternalServerError, err := s.client.Todos.AddOne(
		&todos.AddOneParams{
			Body: &todoclientmodels.Item{
				Completed:   item.Completed,
				Description: swag.String(item.Description),
				ID:          id,
			},
			Context: context.Background(),
		},
	)
	s.Require().NoError(err)
	s.Require().Nil(addOneInternalServerError)
	s.Require().NotNil(addOneCreated)
	s.Require().NotNil(addOneCreated.Payload)

	payload := addOneCreated.Payload

	s.Assert().Equal(item.Description, swag.StringValue(payload.Description))
	s.Assert().Equal(item.Completed, payload.Completed)
	s.Assert().Equal(id, payload.ID)
}

func (s *Suite) TestAddAndUpdateOneHandler() {
	var item Item
	s.fuzzer.Fuzz(&item)

	var id int64
	s.fuzzer.Fuzz(&id)

	addOneCreated, addOneInternalServerError, err := s.client.Todos.AddOne(
		&todos.AddOneParams{
			Body: &todoclientmodels.Item{
				Completed:   item.Completed,
				Description: swag.String(item.Description),
				ID:          id,
			},
			Context: context.Background(),
		},
	)
	s.Require().NoError(err)
	s.Require().Nil(addOneInternalServerError)
	s.Require().NotNil(addOneCreated)
	s.Require().NotNil(addOneCreated.Payload)

	payload := addOneCreated.Payload

	s.Assert().Equal(item.Description, swag.StringValue(payload.Description))
	s.Assert().Equal(item.Completed, payload.Completed)
	s.Assert().Equal(id, payload.ID)

	// update item
	updatedItem := *payload
	s.fuzzer.Fuzz(&updatedItem)

	updatedItem.ID = id

	updateOneOK, updateOneNotFound, updateOneInternalServerError, err := s.client.Todos.UpdateOne(&todos.UpdateOneParams{
		Body: &todoclientmodels.Item{
			Completed:   updatedItem.Completed,
			Description: updatedItem.Description,
			ID:          updatedItem.ID,
		},
		ID:      updatedItem.ID,
		Context: context.Background(),
	})
	s.Require().NoError(err)
	s.Require().Nil(updateOneInternalServerError)
	s.Require().Nil(updateOneNotFound)
	s.Require().NotNil(updateOneOK)
	s.Require().NotNil(updateOneOK.Payload)

	updatedItemResponse := updateOneOK.Payload

	s.Assert().Equal(updatedItem.ID, updatedItemResponse.ID)
	s.Assert().Equal(updatedItem.Completed, updatedItemResponse.Completed)
	s.Assert().EqualValues(updatedItem.Description, updatedItemResponse.Description)
}

func (s *Suite) FindHandler() {
	var item1, item2, item3 Item

	criteria1 := "abc"
	criteria2 := "bcd"
	criteria3 := "def"

	var id1, id2, id3 int64

	// add an item
	{
		s.fuzzer.Fuzz(&item1)
		item1.Description += criteria1

		s.fuzzer.Fuzz(&id1)

		addOneCreated, addOneInternalServerError, err := s.client.Todos.AddOne(
			&todos.AddOneParams{
				Body: &todoclientmodels.Item{
					Completed:   item1.Completed,
					Description: swag.String(item1.Description),
					ID:          id1,
				},
				Context: context.Background(),
			},
		)
		s.Require().NoError(err)
		s.Require().Nil(addOneInternalServerError)
		s.Require().NotNil(addOneCreated)
		s.Require().NotNil(addOneCreated.Payload)

		payload := addOneCreated.Payload

		s.Assert().Equal(item1.Description, swag.StringValue(payload.Description))
		s.Assert().Equal(item1.Completed, payload.Completed)
		s.Assert().Equal(id1, payload.ID)
	}

	// add 2-nd item
	{
		s.fuzzer.Fuzz(&item2)
		item2.Description = item2.Description[:len(item2.Description)/2] + criteria2 + item2.Description[len(item2.Description)/2:]

		s.fuzzer.Fuzz(&id2)

		addOneCreated, addOneInternalServerError, err := s.client.Todos.AddOne(
			&todos.AddOneParams{
				Body: &todoclientmodels.Item{
					Completed:   item2.Completed,
					Description: swag.String(item2.Description),
					ID:          id2,
				},
				Context: context.Background(),
			},
		)
		s.Require().NoError(err)
		s.Require().Nil(addOneInternalServerError)
		s.Require().NotNil(addOneCreated)
		s.Require().NotNil(addOneCreated.Payload)

		payload := addOneCreated.Payload

		s.Assert().Equal(item2.Description, swag.StringValue(payload.Description))
		s.Assert().Equal(item2.Completed, payload.Completed)
		s.Assert().Equal(id2, payload.ID)
	}

	// add 3-rd item
	{
		s.fuzzer.Fuzz(&item3)
		item3.Description = criteria3 + item3.Description

		s.fuzzer.Fuzz(&id3)

		addOneCreated, addOneInternalServerError, err := s.client.Todos.AddOne(
			&todos.AddOneParams{
				Body: &todoclientmodels.Item{
					Completed:   item3.Completed,
					Description: swag.String(item3.Description),
					ID:          id3,
				},
				Context: context.Background(),
			},
		)
		s.Require().NoError(err)
		s.Require().Nil(addOneInternalServerError)
		s.Require().NotNil(addOneCreated)
		s.Require().NotNil(addOneCreated.Payload)

		payload := addOneCreated.Payload

		s.Assert().Equal(item3.Description, swag.StringValue(payload.Description))
		s.Assert().Equal(item3.Completed, payload.Completed)
		s.Assert().Equal(id3, payload.ID)
	}

	{
		findOK, findNotFound, findInternalServerError, err := s.client.Todos.Find(&todos.FindParams{
			XRateLimit: 0,
			Limit:      2,
			Words: []string{
				criteria1,
				criteria2,
			},
			Context: context.Background(),
		})
		s.Require().NoError(err)
		s.Require().Nil(findInternalServerError)
		s.Require().Nil(findNotFound)
		s.Require().NotNil(findOK)
		s.Require().NotNil(findOK.Payload)

		items := findOK.Payload
		s.Assert().Len(items, 2)
		for _, item := range items {
			if item.ID == id1 {
				s.Assert().Equal(item1.Description, swag.StringValue(item.Description))
				s.Assert().Equal(item1.Completed, item.Completed)
			}

			if item.ID == id2 {
				s.Assert().Equal(item2.Description, swag.StringValue(item.Description))
				s.Assert().Equal(item2.Completed, item.Completed)
			}

			if item.ID == id3 {
				s.Assert().Equal(item3.Description, swag.StringValue(item.Description))
				s.Assert().Equal(item3.Completed, item.Completed)
			}
		}
	}
}
