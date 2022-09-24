// Code generated by ent, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// ActivitiesColumns holds the columns for the "activities" table.
	ActivitiesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeString},
		{Name: "name", Type: field.TypeString},
		{Name: "date", Type: field.TypeTime},
		{Name: "area_activities", Type: field.TypeString, Nullable: true},
		{Name: "class_period_activities", Type: field.TypeString, Nullable: true},
	}
	// ActivitiesTable holds the schema information for the "activities" table.
	ActivitiesTable = &schema.Table{
		Name:       "activities",
		Columns:    ActivitiesColumns,
		PrimaryKey: []*schema.Column{ActivitiesColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "activities_areas_activities",
				Columns:    []*schema.Column{ActivitiesColumns[3]},
				RefColumns: []*schema.Column{AreasColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:     "activities_class_periods_activities",
				Columns:    []*schema.Column{ActivitiesColumns[4]},
				RefColumns: []*schema.Column{ClassPeriodsColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// ActivitySyncsColumns holds the columns for the "activity_syncs" table.
	ActivitySyncsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeString},
		{Name: "last_sync_id", Type: field.TypeString},
		{Name: "teacher_activity_syncs", Type: field.TypeString, Nullable: true},
	}
	// ActivitySyncsTable holds the schema information for the "activity_syncs" table.
	ActivitySyncsTable = &schema.Table{
		Name:       "activity_syncs",
		Columns:    ActivitySyncsColumns,
		PrimaryKey: []*schema.Column{ActivitySyncsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "activity_syncs_teachers_activitySyncs",
				Columns:    []*schema.Column{ActivitySyncsColumns[2]},
				RefColumns: []*schema.Column{TeachersColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// AreasColumns holds the columns for the "areas" table.
	AreasColumns = []*schema.Column{
		{Name: "id", Type: field.TypeString},
		{Name: "name", Type: field.TypeString},
		{Name: "points", Type: field.TypeInt},
		{Name: "year_areas", Type: field.TypeString, Nullable: true},
	}
	// AreasTable holds the schema information for the "areas" table.
	AreasTable = &schema.Table{
		Name:       "areas",
		Columns:    AreasColumns,
		PrimaryKey: []*schema.Column{AreasColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "areas_years_areas",
				Columns:    []*schema.Column{AreasColumns[3]},
				RefColumns: []*schema.Column{YearsColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// AttendancesColumns holds the columns for the "attendances" table.
	AttendancesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeString},
		{Name: "value", Type: field.TypeEnum, Enums: []string{"Presente", "Falta", "Atraso", "Licencia"}},
		{Name: "attendance_day_attendances", Type: field.TypeString, Nullable: true},
		{Name: "student_attendances", Type: field.TypeString, Nullable: true},
	}
	// AttendancesTable holds the schema information for the "attendances" table.
	AttendancesTable = &schema.Table{
		Name:       "attendances",
		Columns:    AttendancesColumns,
		PrimaryKey: []*schema.Column{AttendancesColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "attendances_attendance_days_attendances",
				Columns:    []*schema.Column{AttendancesColumns[2]},
				RefColumns: []*schema.Column{AttendanceDaysColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:     "attendances_students_attendances",
				Columns:    []*schema.Column{AttendancesColumns[3]},
				RefColumns: []*schema.Column{StudentsColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// AttendanceDaysColumns holds the columns for the "attendance_days" table.
	AttendanceDaysColumns = []*schema.Column{
		{Name: "id", Type: field.TypeString},
		{Name: "day", Type: field.TypeTime},
		{Name: "class_period_attendance_days", Type: field.TypeString, Nullable: true},
	}
	// AttendanceDaysTable holds the schema information for the "attendance_days" table.
	AttendanceDaysTable = &schema.Table{
		Name:       "attendance_days",
		Columns:    AttendanceDaysColumns,
		PrimaryKey: []*schema.Column{AttendanceDaysColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "attendance_days_class_periods_attendanceDays",
				Columns:    []*schema.Column{AttendanceDaysColumns[2]},
				RefColumns: []*schema.Column{ClassPeriodsColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// AttendanceDaySyncsColumns holds the columns for the "attendance_day_syncs" table.
	AttendanceDaySyncsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeString},
		{Name: "last_sync_id", Type: field.TypeString},
		{Name: "teacher_attendance_day_syncs", Type: field.TypeString, Nullable: true},
	}
	// AttendanceDaySyncsTable holds the schema information for the "attendance_day_syncs" table.
	AttendanceDaySyncsTable = &schema.Table{
		Name:       "attendance_day_syncs",
		Columns:    AttendanceDaySyncsColumns,
		PrimaryKey: []*schema.Column{AttendanceDaySyncsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "attendance_day_syncs_teachers_attendanceDaySyncs",
				Columns:    []*schema.Column{AttendanceDaySyncsColumns[2]},
				RefColumns: []*schema.Column{TeachersColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// AttendanceSyncsColumns holds the columns for the "attendance_syncs" table.
	AttendanceSyncsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeString},
		{Name: "last_sync_id", Type: field.TypeString},
		{Name: "teacher_attendance_syncs", Type: field.TypeString, Nullable: true},
	}
	// AttendanceSyncsTable holds the schema information for the "attendance_syncs" table.
	AttendanceSyncsTable = &schema.Table{
		Name:       "attendance_syncs",
		Columns:    AttendanceSyncsColumns,
		PrimaryKey: []*schema.Column{AttendanceSyncsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "attendance_syncs_teachers_attendanceSyncs",
				Columns:    []*schema.Column{AttendanceSyncsColumns[2]},
				RefColumns: []*schema.Column{TeachersColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// ClassesColumns holds the columns for the "classes" table.
	ClassesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeString},
		{Name: "parallel", Type: field.TypeString},
		{Name: "grade_classes", Type: field.TypeString, Nullable: true},
		{Name: "school_classes", Type: field.TypeString, Nullable: true},
		{Name: "subject_classes", Type: field.TypeString, Nullable: true},
		{Name: "teacher_classes", Type: field.TypeString, Nullable: true},
		{Name: "year_classes", Type: field.TypeString, Nullable: true},
	}
	// ClassesTable holds the schema information for the "classes" table.
	ClassesTable = &schema.Table{
		Name:       "classes",
		Columns:    ClassesColumns,
		PrimaryKey: []*schema.Column{ClassesColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "classes_grades_classes",
				Columns:    []*schema.Column{ClassesColumns[2]},
				RefColumns: []*schema.Column{GradesColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:     "classes_schools_classes",
				Columns:    []*schema.Column{ClassesColumns[3]},
				RefColumns: []*schema.Column{SchoolsColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:     "classes_subjects_classes",
				Columns:    []*schema.Column{ClassesColumns[4]},
				RefColumns: []*schema.Column{SubjectsColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:     "classes_teachers_classes",
				Columns:    []*schema.Column{ClassesColumns[5]},
				RefColumns: []*schema.Column{TeachersColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:     "classes_years_classes",
				Columns:    []*schema.Column{ClassesColumns[6]},
				RefColumns: []*schema.Column{YearsColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// ClassPeriodsColumns holds the columns for the "class_periods" table.
	ClassPeriodsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeString},
		{Name: "start", Type: field.TypeTime},
		{Name: "end", Type: field.TypeTime},
		{Name: "finished", Type: field.TypeBool},
		{Name: "class_class_periods", Type: field.TypeString, Nullable: true},
		{Name: "period_class_periods", Type: field.TypeString, Nullable: true},
	}
	// ClassPeriodsTable holds the schema information for the "class_periods" table.
	ClassPeriodsTable = &schema.Table{
		Name:       "class_periods",
		Columns:    ClassPeriodsColumns,
		PrimaryKey: []*schema.Column{ClassPeriodsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "class_periods_classes_classPeriods",
				Columns:    []*schema.Column{ClassPeriodsColumns[4]},
				RefColumns: []*schema.Column{ClassesColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:     "class_periods_periods_classPeriods",
				Columns:    []*schema.Column{ClassPeriodsColumns[5]},
				RefColumns: []*schema.Column{PeriodsColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// ClassPeriodSyncsColumns holds the columns for the "class_period_syncs" table.
	ClassPeriodSyncsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeString},
		{Name: "last_sync_id", Type: field.TypeString},
		{Name: "teacher_class_period_syncs", Type: field.TypeString, Nullable: true},
	}
	// ClassPeriodSyncsTable holds the schema information for the "class_period_syncs" table.
	ClassPeriodSyncsTable = &schema.Table{
		Name:       "class_period_syncs",
		Columns:    ClassPeriodSyncsColumns,
		PrimaryKey: []*schema.Column{ClassPeriodSyncsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "class_period_syncs_teachers_classPeriodSyncs",
				Columns:    []*schema.Column{ClassPeriodSyncsColumns[2]},
				RefColumns: []*schema.Column{TeachersColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// DptosColumns holds the columns for the "dptos" table.
	DptosColumns = []*schema.Column{
		{Name: "id", Type: field.TypeString},
		{Name: "name", Type: field.TypeString},
	}
	// DptosTable holds the schema information for the "dptos" table.
	DptosTable = &schema.Table{
		Name:       "dptos",
		Columns:    DptosColumns,
		PrimaryKey: []*schema.Column{DptosColumns[0]},
	}
	// GradesColumns holds the columns for the "grades" table.
	GradesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeString},
		{Name: "name", Type: field.TypeString},
	}
	// GradesTable holds the schema information for the "grades" table.
	GradesTable = &schema.Table{
		Name:       "grades",
		Columns:    GradesColumns,
		PrimaryKey: []*schema.Column{GradesColumns[0]},
	}
	// MunicipiosColumns holds the columns for the "municipios" table.
	MunicipiosColumns = []*schema.Column{
		{Name: "id", Type: field.TypeString},
		{Name: "name", Type: field.TypeString},
		{Name: "provincia_municipios", Type: field.TypeString, Nullable: true},
	}
	// MunicipiosTable holds the schema information for the "municipios" table.
	MunicipiosTable = &schema.Table{
		Name:       "municipios",
		Columns:    MunicipiosColumns,
		PrimaryKey: []*schema.Column{MunicipiosColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "municipios_provincia_municipios",
				Columns:    []*schema.Column{MunicipiosColumns[2]},
				RefColumns: []*schema.Column{ProvinciaColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// PeriodsColumns holds the columns for the "periods" table.
	PeriodsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeString},
		{Name: "name", Type: field.TypeString},
		{Name: "year_periods", Type: field.TypeString, Nullable: true},
	}
	// PeriodsTable holds the schema information for the "periods" table.
	PeriodsTable = &schema.Table{
		Name:       "periods",
		Columns:    PeriodsColumns,
		PrimaryKey: []*schema.Column{PeriodsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "periods_years_periods",
				Columns:    []*schema.Column{PeriodsColumns[2]},
				RefColumns: []*schema.Column{YearsColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// ProvinciaColumns holds the columns for the "provincia" table.
	ProvinciaColumns = []*schema.Column{
		{Name: "id", Type: field.TypeString},
		{Name: "name", Type: field.TypeString},
		{Name: "dpto_provincias", Type: field.TypeString, Nullable: true},
	}
	// ProvinciaTable holds the schema information for the "provincia" table.
	ProvinciaTable = &schema.Table{
		Name:       "provincia",
		Columns:    ProvinciaColumns,
		PrimaryKey: []*schema.Column{ProvinciaColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "provincia_dptos_provincias",
				Columns:    []*schema.Column{ProvinciaColumns[2]},
				RefColumns: []*schema.Column{DptosColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// SchoolsColumns holds the columns for the "schools" table.
	SchoolsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeString},
		{Name: "name", Type: field.TypeString},
		{Name: "lat", Type: field.TypeString},
		{Name: "lon", Type: field.TypeString},
		{Name: "municipio_schools", Type: field.TypeString, Nullable: true},
	}
	// SchoolsTable holds the schema information for the "schools" table.
	SchoolsTable = &schema.Table{
		Name:       "schools",
		Columns:    SchoolsColumns,
		PrimaryKey: []*schema.Column{SchoolsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "schools_municipios_schools",
				Columns:    []*schema.Column{SchoolsColumns[4]},
				RefColumns: []*schema.Column{MunicipiosColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// ScoresColumns holds the columns for the "scores" table.
	ScoresColumns = []*schema.Column{
		{Name: "id", Type: field.TypeString},
		{Name: "points", Type: field.TypeInt},
		{Name: "activity_scores", Type: field.TypeString, Nullable: true},
		{Name: "student_scores", Type: field.TypeString, Nullable: true},
	}
	// ScoresTable holds the schema information for the "scores" table.
	ScoresTable = &schema.Table{
		Name:       "scores",
		Columns:    ScoresColumns,
		PrimaryKey: []*schema.Column{ScoresColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "scores_activities_scores",
				Columns:    []*schema.Column{ScoresColumns[2]},
				RefColumns: []*schema.Column{ActivitiesColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:     "scores_students_scores",
				Columns:    []*schema.Column{ScoresColumns[3]},
				RefColumns: []*schema.Column{StudentsColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// ScoreSyncsColumns holds the columns for the "score_syncs" table.
	ScoreSyncsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeString},
		{Name: "last_sync_id", Type: field.TypeString},
		{Name: "teacher_score_syncs", Type: field.TypeString, Nullable: true},
	}
	// ScoreSyncsTable holds the schema information for the "score_syncs" table.
	ScoreSyncsTable = &schema.Table{
		Name:       "score_syncs",
		Columns:    ScoreSyncsColumns,
		PrimaryKey: []*schema.Column{ScoreSyncsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "score_syncs_teachers_scoreSyncs",
				Columns:    []*schema.Column{ScoreSyncsColumns[2]},
				RefColumns: []*schema.Column{TeachersColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// StudentsColumns holds the columns for the "students" table.
	StudentsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeString},
		{Name: "name", Type: field.TypeString},
		{Name: "last_name", Type: field.TypeString},
		{Name: "ci", Type: field.TypeString},
		{Name: "class_students", Type: field.TypeString, Nullable: true},
	}
	// StudentsTable holds the schema information for the "students" table.
	StudentsTable = &schema.Table{
		Name:       "students",
		Columns:    StudentsColumns,
		PrimaryKey: []*schema.Column{StudentsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "students_classes_students",
				Columns:    []*schema.Column{StudentsColumns[4]},
				RefColumns: []*schema.Column{ClassesColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// StudentSyncsColumns holds the columns for the "student_syncs" table.
	StudentSyncsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeString},
		{Name: "last_sync_id", Type: field.TypeString},
		{Name: "teacher_student_syncs", Type: field.TypeString, Nullable: true},
	}
	// StudentSyncsTable holds the schema information for the "student_syncs" table.
	StudentSyncsTable = &schema.Table{
		Name:       "student_syncs",
		Columns:    StudentSyncsColumns,
		PrimaryKey: []*schema.Column{StudentSyncsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "student_syncs_teachers_studentSyncs",
				Columns:    []*schema.Column{StudentSyncsColumns[2]},
				RefColumns: []*schema.Column{TeachersColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// SubjectsColumns holds the columns for the "subjects" table.
	SubjectsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeString},
		{Name: "name", Type: field.TypeString},
	}
	// SubjectsTable holds the schema information for the "subjects" table.
	SubjectsTable = &schema.Table{
		Name:       "subjects",
		Columns:    SubjectsColumns,
		PrimaryKey: []*schema.Column{SubjectsColumns[0]},
	}
	// TeachersColumns holds the columns for the "teachers" table.
	TeachersColumns = []*schema.Column{
		{Name: "id", Type: field.TypeString},
		{Name: "name", Type: field.TypeString},
		{Name: "last_name", Type: field.TypeString},
		{Name: "email", Type: field.TypeString},
		{Name: "password", Type: field.TypeString},
	}
	// TeachersTable holds the schema information for the "teachers" table.
	TeachersTable = &schema.Table{
		Name:       "teachers",
		Columns:    TeachersColumns,
		PrimaryKey: []*schema.Column{TeachersColumns[0]},
	}
	// YearsColumns holds the columns for the "years" table.
	YearsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeString},
		{Name: "value", Type: field.TypeInt},
	}
	// YearsTable holds the schema information for the "years" table.
	YearsTable = &schema.Table{
		Name:       "years",
		Columns:    YearsColumns,
		PrimaryKey: []*schema.Column{YearsColumns[0]},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		ActivitiesTable,
		ActivitySyncsTable,
		AreasTable,
		AttendancesTable,
		AttendanceDaysTable,
		AttendanceDaySyncsTable,
		AttendanceSyncsTable,
		ClassesTable,
		ClassPeriodsTable,
		ClassPeriodSyncsTable,
		DptosTable,
		GradesTable,
		MunicipiosTable,
		PeriodsTable,
		ProvinciaTable,
		SchoolsTable,
		ScoresTable,
		ScoreSyncsTable,
		StudentsTable,
		StudentSyncsTable,
		SubjectsTable,
		TeachersTable,
		YearsTable,
	}
)

func init() {
	ActivitiesTable.ForeignKeys[0].RefTable = AreasTable
	ActivitiesTable.ForeignKeys[1].RefTable = ClassPeriodsTable
	ActivitySyncsTable.ForeignKeys[0].RefTable = TeachersTable
	AreasTable.ForeignKeys[0].RefTable = YearsTable
	AttendancesTable.ForeignKeys[0].RefTable = AttendanceDaysTable
	AttendancesTable.ForeignKeys[1].RefTable = StudentsTable
	AttendanceDaysTable.ForeignKeys[0].RefTable = ClassPeriodsTable
	AttendanceDaySyncsTable.ForeignKeys[0].RefTable = TeachersTable
	AttendanceSyncsTable.ForeignKeys[0].RefTable = TeachersTable
	ClassesTable.ForeignKeys[0].RefTable = GradesTable
	ClassesTable.ForeignKeys[1].RefTable = SchoolsTable
	ClassesTable.ForeignKeys[2].RefTable = SubjectsTable
	ClassesTable.ForeignKeys[3].RefTable = TeachersTable
	ClassesTable.ForeignKeys[4].RefTable = YearsTable
	ClassPeriodsTable.ForeignKeys[0].RefTable = ClassesTable
	ClassPeriodsTable.ForeignKeys[1].RefTable = PeriodsTable
	ClassPeriodSyncsTable.ForeignKeys[0].RefTable = TeachersTable
	MunicipiosTable.ForeignKeys[0].RefTable = ProvinciaTable
	PeriodsTable.ForeignKeys[0].RefTable = YearsTable
	ProvinciaTable.ForeignKeys[0].RefTable = DptosTable
	SchoolsTable.ForeignKeys[0].RefTable = MunicipiosTable
	ScoresTable.ForeignKeys[0].RefTable = ActivitiesTable
	ScoresTable.ForeignKeys[1].RefTable = StudentsTable
	ScoreSyncsTable.ForeignKeys[0].RefTable = TeachersTable
	StudentsTable.ForeignKeys[0].RefTable = ClassesTable
	StudentSyncsTable.ForeignKeys[0].RefTable = TeachersTable
}
