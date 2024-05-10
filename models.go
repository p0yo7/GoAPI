// models.go

package GoAPI

import (
	"gorm.io/gorm"
	"time"
)

type User struct {
	gorm.Model
	id_user uint "gorm:primaryKey"
	username string "gorm:unique"
	password string
	name string
	contact_phone string
	contact_email string
	user_role string
}

type exercises_plan struct {
	gorm.Model
	id_routine uint "gorm:primaryKey"
	id_user int
	User User "gorm:foreignKey:id_user"
	sets_completed int
	sets_planned int
	routine_date time.time
}

type exercise struct {
	gorm.Model
	id_exercise int "gorm:primaryKey"
	name string
	description string
	instructions string
	image_name string
}

type routine_exercise struct {
	gorm.Model
	id_routine_exercise uint "gorm:primaryKey"
	id_routine uint 
	exercises_plan exercises_plan "gorm:foreignKey:id_routine"
	id_exercise uint
	exercise exercise "gorm:foreignKey:id_exercise"
	reps_planned int
	reps_completed int
}
