package graph_test

import (
	"github.com/99designs/gqlgen/client"
	"github.com/99designs/gqlgen/graphql/handler"
	generated "github.com/AlexandrLitkevich/pet-trello/graph/generated"
	mock2 "github.com/AlexandrLitkevich/pet-trello/graph/mock"
	"github.com/AlexandrLitkevich/pet-trello/graph/model"
	resolvers "github.com/AlexandrLitkevich/pet-trello/graph/resolvers"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"testing"
)

var (
	id   = "1"
	name = "one"
)

func TestQueryResolver_User(t *testing.T) {
	t.Run("Start test", func(t *testing.T) {
		// implement resolver
		testUserSevice := new(mock2.UserServiceMock)
		// creteStruct
		resolver := resolvers.Resolver{
			UserService: testUserSevice,
		}
		// Create new server
		c := client.New(handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &resolver})))
		u := model.User{
			ID:   id,
			Name: name,
		}
		// это мок
		testUserSevice.On("GetUserById", mock.AnythingOfType("string")).Return(&u)
		t.Log(c)

		var resp struct {
			User struct {
				ID, Name string
			}
		}
		t.Log(resp)
		// request
		q := `
		query GetUser($id: String!) {
			user(id: $id) {
				  id
                  name
			}
		}
		`
		// client request
		err := c.Post(q, &resp, client.Var("id", "1"))
		if err != nil {
			t.Log(err)
		}
		testUserSevice.AssertCalled(t, "GetUserById", "1")

		require.Equal(t, "1", resp.User.ID)
		require.Equal(t, "one", resp.User.Name)

	})
}
