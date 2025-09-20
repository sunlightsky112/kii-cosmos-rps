package rpsKeeper

import (
	"context"

	"challenge/x/rps/types"
)

// SetStudent stores a student in the Students map
func (k Keeper) SetStudent(ctx context.Context, student types.Student) error {
	return k.Students.Set(ctx, student.Id, student)
}

// GetStudent retrieves a student by ID
func (k Keeper) GetStudent(ctx context.Context, id string) (types.Student, error) {
	return k.Students.Get(ctx, id)
}

// DeleteStudent removes a student by ID
func (k Keeper) DeleteStudent(ctx context.Context, id string) error {
	return k.Students.Remove(ctx, id)
}

// GetAllStudents returns all students
func (k Keeper) GetAllStudents(ctx context.Context) ([]types.Student, error) {
	var students []types.Student

	// Iterate through all students
	err := k.Students.Walk(ctx, nil, func(key string, student types.Student) (stop bool, err error) {
		students = append(students, student)
		return false, nil
	})

	return students, err
}
