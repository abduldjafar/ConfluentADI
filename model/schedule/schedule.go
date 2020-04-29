package schedule

type ReqJSONScheduleV1 struct {
	JSONScheduleV1 struct {
		CIFBankHolidayRunning interface{} `json:"CIF_bank_holiday_running"`
		CIFStpIndicator       string      `json:"CIF_stp_indicator"`
		CIFTrainUID           string      `json:"CIF_train_uid"`
		ApplicableTimetable   string      `json:"applicable_timetable"`
		AtocCode              string      `json:"atoc_code"`
		NewScheduleSegment    struct {
			TractionClass string `json:"traction_class"`
			UicCode       string `json:"uic_code"`
		} `json:"new_schedule_segment"`
		ScheduleDaysRuns string `json:"schedule_days_runs"`
		ScheduleEndDate  string `json:"schedule_end_date"`
		ScheduleSegment  struct {
			SignallingID                string      `json:"signalling_id"`
			CIFTrainCategory            string      `json:"CIF_train_category"`
			CIFHeadcode                 string      `json:"CIF_headcode"`
			CIFCourseIndicator          int         `json:"CIF_course_indicator"`
			CIFTrainServiceCode         string      `json:"CIF_train_service_code"`
			CIFBusinessSector           string      `json:"CIF_business_sector"`
			CIFPowerType                string      `json:"CIF_power_type"`
			CIFTimingLoad               interface{} `json:"CIF_timing_load"`
			CIFSpeed                    string      `json:"CIF_speed"`
			CIFOperatingCharacteristics interface{} `json:"CIF_operating_characteristics"`
			CIFTrainClass               interface{} `json:"CIF_train_class"`
			CIFSleepers                 interface{} `json:"CIF_sleepers"`
			CIFReservations             interface{} `json:"CIF_reservations"`
			CIFConnectionIndicator      interface{} `json:"CIF_connection_indicator"`
			CIFCateringCode             interface{} `json:"CIF_catering_code"`
			CIFServiceBranding          string      `json:"CIF_service_branding"`
			ScheduleLocation            []struct {
				LocationType         string      `json:"location_type"`
				RecordIdentity       string      `json:"record_identity"`
				TiplocCode           string      `json:"tiploc_code"`
				TiplocInstance       interface{} `json:"tiploc_instance"`
				Departure            string      `json:"departure,omitempty"`
				PublicDeparture      interface{} `json:"public_departure,omitempty"`
				Platform             string      `json:"platform"`
				Line                 string      `json:"line,omitempty"`
				EngineeringAllowance interface{} `json:"engineering_allowance,omitempty"`
				PathingAllowance     interface{} `json:"pathing_allowance,omitempty"`
				PerformanceAllowance interface{} `json:"performance_allowance,omitempty"`
				Arrival              interface{} `json:"arrival,omitempty"`
				Pass                 string      `json:"pass,omitempty"`
				PublicArrival        interface{} `json:"public_arrival,omitempty"`
				Path                 interface{} `json:"path,omitempty"`
			} `json:"schedule_location"`
		} `json:"schedule_segment"`
		ScheduleStartDate string `json:"schedule_start_date"`
		TrainStatus       string `json:"train_status"`
		TransactionType   string `json:"transaction_type"`
	} `json:"JsonScheduleV1"`
}
