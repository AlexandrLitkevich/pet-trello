package graph

import (
	"github.com/99designs/gqlgen/client"
	"github.com/99designs/gqlgen/graphql/handler"
	generated "github.com/AlexandrLitkevich/pet-trello/graph/generated"
	mock2 "github.com/AlexandrLitkevich/pet-trello/graph/mock"
	"github.com/AlexandrLitkevich/pet-trello/graph/model"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"testing"
)

var (
	id   = "1"
	name = "one"
)

func TestQueryResolver_User(t *testing.T) {
	t.Run("Start test queryresolver user", func(t *testing.T) {
		testUserSevice := mock2.NewUserServiceMock()
		resolver := Resolver{
			UserService: testUserSevice,
		}
		c := client.New(handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &resolver})))
		u := model.User{
			ID:   id,
			Name: name,
		}
		testUserSevice.On("GetUserById", mock.AnythingOfType("string")).Return(&u)

		var resp struct {
			User struct {
				ID, Name string
			}
		}

		q := `
		  query GetUser(id: String!) { 
			user(id: $id) { 
			  id, name
			} 
		  }
    	`
		c.MustPost(q, &resp, client.Var("id", "1"))
		testUserSevice.AssertCalled(t, "GetUserById", "1")

		require.Equal(t, "1", resp.User.ID)
		require.Equal(t, "1", resp.User.Name)

	})
}

/*
func TestQueryResolver_User(t *testing.T) {
	t.Run("should query user correctly", func(t *testing.T) {
		testUserService := new(mocks.MockedUserService)
		resolvers := resolver.Resolver{UserService: testUserService}
		c := client.New(handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &resolvers})))
		u := model.UserDetail{User: model.User{Loginname: &loginname, AvatarURL: &avatarURL}, Score: &score, CreateAt: &createAt}
		testUserService.On("GetUserByLoginname", mock.AnythingOfType("string")).Return(&u)
		var resp struct {
			User struct {
				Loginname, AvatarURL, CreateAt string
				Score                          int
			}
		}
		q := `
      query GetUser($loginname: String!) {
        user(loginname: $loginname) {
          loginname
          avatarUrl
          createAt
          score
        }
      }
    `
		c.MustPost(q, &resp, client.Var("loginname", "mrdulin"))
		testUserService.AssertCalled(t, "GetUserByLoginname", "mrdulin")

		require.Equal(t, "mrdulin", resp.User.Loginname)
		require.Equal(t, "avatar.jpg", resp.User.AvatarURL)
		require.Equal(t, 50, resp.User.Score)
		require.Equal(t, "1900-01-01", resp.User.CreateAt)
	})
}


*/
