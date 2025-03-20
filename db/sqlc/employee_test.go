package db

import (
	"GoBrewAPI/util"
	"context"
	"database/sql"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func createRandomEmployee(t *testing.T) Employee {
	var roleQ util.RoleENUM = util.SWE

	arg := CreateEmployeeParams{
		Firstname: util.RandomFirstName(),
		Lastname:  util.RandomLastName(),
		Password:  util.PwdGen(),
		Role:      roleQ.String(),
		CreatedAt: util.RandomDate(),
		IsAdmin:   true,
	}

	employee, err := testQueries.CreateEmployee(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, employee)

	require.Equal(t, arg.Firstname, employee.Firstname)
	require.Equal(t, arg.Lastname, employee.Lastname)
	require.Equal(t, arg.Password, employee.Password)
	require.Equal(t, arg.Role, employee.Role)
	require.Equal(t, arg.CreatedAt.Time.In(time.UTC), employee.CreatedAt.Time.In(time.UTC))
	require.Equal(t, arg.IsAdmin, employee.IsAdmin)

	require.NotZero(t, employee.ID)

	return employee
}

func deleteIDEmployee(employee Employee) error {
	err := testQueries.DeleteEmployee(context.Background(), employee.ID)
	return err
}

func TestCreateEmployee(t *testing.T) {
	employee := createRandomEmployee(t)
	deleteIDEmployee(employee)
}

func TestGetEmployee(t *testing.T) {
	employee := createRandomEmployee(t)
	employeeQ, err := testQueries.GetEmployee(context.Background(), employee.ID)
	require.NoError(t, err)
	require.NotEmpty(t, employeeQ)

	require.Equal(t, employeeQ.Firstname, employee.Firstname)
	require.Equal(t, employeeQ.Lastname, employee.Lastname)
	require.Equal(t, employeeQ.Password, employee.Password)
	require.Equal(t, employeeQ.Role, employee.Role)
	require.Equal(t, employeeQ.CreatedAt.Time.In(time.UTC), employee.CreatedAt.Time.In(time.UTC))
	require.Equal(t, employeeQ.IsAdmin, employee.IsAdmin)
	
	deleteIDEmployee(employeeQ)
}

func TestUpdateEmployee(t *testing.T) {
	employee := createRandomEmployee(t)
	var roleQ util.RoleENUM = util.SWE

	arg := UpdateEmployeeParams{
		Firstname: util.RandomFirstName(),
		Lastname:  util.RandomLastName(),
		Password:  util.PwdGen(),
		Role:      roleQ.String(),
		CreatedAt: util.RandomDate(),
		IsAdmin:   true,
		ID:       employee.ID,
	}

	employeeQ, err := testQueries.UpdateEmployee(context.Background(), arg)

	require.NoError(t, err)
	require.NotEmpty(t, employeeQ)
	require.Equal(t, employeeQ.Firstname, arg.Firstname)
	require.Equal(t, employeeQ.Lastname, arg.Lastname)
	require.Equal(t, employeeQ.Password, arg.Password)
	require.Equal(t, employeeQ.Role, arg.Role)
	require.Equal(t, employeeQ.CreatedAt.Time.In(time.UTC), arg.CreatedAt.Time.In(time.UTC))
	require.Equal(t, employeeQ.IsAdmin, arg.IsAdmin)

	deleteIDEmployee(employeeQ)
}

func TestDeleteEmployee(t *testing.T) {
	employee := createRandomEmployee(t)
	err := deleteIDEmployee(employee)
	require.NoError(t, err)

	employeeQ, err := testQueries.GetEmployee(context.Background(), employee.ID)
	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, employeeQ)
}
