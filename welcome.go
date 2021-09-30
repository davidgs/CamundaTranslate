package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"path/filepath"
)

type WelcomeInterface struct {
	Labels WelcomeLabels `json:"labels"`
}
type WelcomeLabels struct {
		AppVendor                                   string `json:"APP_VENDOR"`
		AdminApp                                    string `json:"ADMIN_APP"`
		TasklistApp                                 string `json:"TASKLIST_APP"`
		CockpitApp                                  string `json:"COCKPIT_APP"`
		ToggleNavigation                            string `json:"TOGGLE_NAVIGATION"`
		Submit                                      string `json:"SUBMIT"`
		Save                                        string `json:"SAVE"`
		Start                                       string `json:"START"`
		Abort                                       string `json:"ABORT"`
		Close                                       string `json:"CLOSE"`
		Cancel                                      string `json:"CANCEL"`
		Back                                        string `json:"BACK"`
		Delete                                      string `json:"DELETE"`
		Set                                         string `json:"SET"`
		Change                                      string `json:"CHANGE"`
		Task                                        string `json:"TASK"`
		Filter                                      string `json:"FILTER"`
		Tasks                                       string `json:"TASKS"`
		Filters                                     string `json:"FILTERS"`
		Loading                                     string `json:"LOADING"`
		RequiredField                               string `json:"REQUIRED_FIELD"`
		RequiredIntegerField                        string `json:"REQUIRED_INTEGER_FIELD"`
		RequiredHexColorField                       string `json:"REQUIRED_HEX_COLOR_FIELD"`
		ProcessDefinition                           string `json:"PROCESS_DEFINITION"`
		ProcessInstance                             string `json:"PROCESS_INSTANCE"`
		ManageAccount                               string `json:"MANAGE_ACCOUNT"`
		Asc                                         string `json:"ASC"`
		Desc                                        string `json:"DESC"`
		Boolean                                     string `json:"BOOLEAN"`
		Short                                       string `json:"SHORT"`
		Integer                                     string `json:"INTEGER"`
		Long                                        string `json:"LONG"`
		Double                                      string `json:"DOUBLE"`
		Date                                        string `json:"DATE"`
		String                                      string `json:"STRING"`
		AuthLogoutSuccessful                        string `json:"AUTH_LOGOUT_SUCCESSFUL"`
		AuthLogoutThanks                            string `json:"AUTH_LOGOUT_THANKS"`
		AuthDayContextMorning                       string `json:"AUTH_DAY_CONTEXT_MORNING"`
		AuthDayContextDay                           string `json:"AUTH_DAY_CONTEXT_DAY"`
		AuthDayContextAfternoon                     string `json:"AUTH_DAY_CONTEXT_AFTERNOON"`
		AuthDayContextEvening                       string `json:"AUTH_DAY_CONTEXT_EVENING"`
		AuthDayContextNight                         string `json:"AUTH_DAY_CONTEXT_NIGHT"`
		AuthDayContextWeekend                       string `json:"AUTH_DAY_CONTEXT_WEEKEND"`
		AuthFailedToDisplayResource                 string `json:"AUTH_FAILED_TO_DISPLAY_RESOURCE"`
		AuthAuthenticationFailed                    string `json:"AUTH_AUTHENTICATION_FAILED"`
		PageLoginErrorMsg                           string `json:"PAGE_LOGIN_ERROR_MSG"`
		PageLoginUsername                           string `json:"PAGE_LOGIN_USERNAME"`
		PageLoginPassword                           string `json:"PAGE_LOGIN_PASSWORD"`
		PageLoginPleaseSignIn                       string `json:"PAGE_LOGIN_PLEASE_SIGN_IN"`
		PageLoginSignInAction                       string `json:"PAGE_LOGIN_SIGN_IN_ACTION"`
		PageLoginFailed                             string `json:"PAGE_LOGIN_FAILED"`
		SignOutAction                               string `json:"SIGN_OUT_ACTION"`
		MyProfile                                   string `json:"MY_PROFILE"`
		DirectiveEngineSelectTooltip                string `json:"DIRECTIVE_ENGINE_SELECT_TOOLTIP"`
		DirectiveEngineSelectStatusNotFound         string `json:"DIRECTIVE_ENGINE_SELECT_STATUS_NOT_FOUND"`
		DirectiveEngineSelectMessageNotFound        string `json:"DIRECTIVE_ENGINE_SELECT_MESSAGE_NOT_FOUND"`
		DirectiveInplaceTextfieldErrorMsg           string `json:"DIRECTIVE_INPLACE_TEXTFIELD_ERROR_MSG"`
		PagesStatusServerError                      string `json:"PAGES_STATUS_SERVER_ERROR"`
		PagesMsgServerError                         string `json:"PAGES_MSG_SERVER_ERROR"`
		PagesStatusRequestTimeout                   string `json:"PAGES_STATUS_REQUEST_TIMEOUT"`
		PagesMsgRequestTimeout                      string `json:"PAGES_MSG_REQUEST_TIMEOUT"`
		PagesStatusAccessDenied                     string `json:"PAGES_STATUS_ACCESS_DENIED"`
		PagesMsgAccessDenied                        string `json:"PAGES_MSG_ACCESS_DENIED"`
		PagesMsgAccessDeniedResourceID              string `json:"PAGES_MSG_ACCESS_DENIED_RESOURCE_ID"`
		PagesMsgActionDenied                        string `json:"PAGES_MSG_ACTION_DENIED"`
		PagesStatusNotFound                         string `json:"PAGES_STATUS_NOT_FOUND"`
		PagesMsgNotFound                            string `json:"PAGES_MSG_NOT_FOUND"`
		PagesStatusCommunicationError               string `json:"PAGES_STATUS_COMMUNICATION_ERROR"`
		PagesMsgCommunicationError                  string `json:"PAGES_MSG_COMMUNICATION_ERROR"`
		PagesMsgEngineNotExists                     string `json:"PAGES_MSG_ENGINE_NOT_EXISTS"`
		ServicesResourceResolverIDNotFound          string `json:"SERVICES_RESOURCE_RESOLVER_ID_NOT_FOUND"`
		ServicesResourceResolverAuthFailed          string `json:"SERVICES_RESOURCE_RESOLVER_AUTH_FAILED"`
		ServicesResourceResolverReceivedStatus      string `json:"SERVICES_RESOURCE_RESOLVER_RECEIVED_STATUS"`
		ServicesResourceResolverDisplayFailed       string `json:"SERVICES_RESOURCE_RESOLVER_DISPLAY_FAILED"`
		ServicesResourceResolverNoPromise           string `json:"SERVICES_RESOURCE_RESOLVER_NO_PROMISE"`
		CamWidgetBpmnViewerError                    string `json:"CAM_WIDGET_BPMN_VIEWER_ERROR"`
		CamWidgetBpmnViewerLoading                  string `json:"CAM_WIDGET_BPMN_VIEWER_LOADING"`
		CamWidgetCopy                               string `json:"CAM_WIDGET_COPY"`
		CamWidgetCopyLink                           string `json:"CAM_WIDGET_COPY_LINK"`
		CamWidgetCmmnViewerError                    string `json:"CAM_WIDGET_CMMN_VIEWER_ERROR"`
		CamWidgetCmmnViewerLoading                  string `json:"CAM_WIDGET_CMMN_VIEWER_LOADING"`
		CamWidgetDmnViewerError                     string `json:"CAM_WIDGET_DMN_VIEWER_ERROR"`
		CamWidgetDmnViewerLoading                   string `json:"CAM_WIDGET_DMN_VIEWER_LOADING"`
		CamWidgetFooterTimezone                     string `json:"CAM_WIDGET_FOOTER_TIMEZONE"`
		CamWidgetFooterPoweredBy                    string `json:"CAM_WIDGET_FOOTER_POWERED_BY"`
		CamWidgetHeaderToggleNavigation             string `json:"CAM_WIDGET_HEADER_TOGGLE_NAVIGATION"`
		CamWidgetHeaderMyProfile                    string `json:"CAM_WIDGET_HEADER_MY_PROFILE"`
		CamWidgetHeaderSignOut                      string `json:"CAM_WIDGET_HEADER_SIGN_OUT"`
		CamWidgetHeaderSmallScreenWarning           string `json:"CAM_WIDGET_HEADER_SMALL_SCREEN_WARNING"`
		CamWidgetSearchMatchType                    string `json:"CAM_WIDGET_SEARCH_MATCH_TYPE"`
		CamWidgetSearchMatchTypeAny                 string `json:"CAM_WIDGET_SEARCH_MATCH_TYPE_ANY"`
		CamWidgetSearchMatchTypeAll                 string `json:"CAM_WIDGET_SEARCH_MATCH_TYPE_ALL"`
		CamWidgetSearchTotalNumberResults           string `json:"CAM_WIDGET_SEARCH_TOTAL_NUMBER_RESULTS"`
		CamWidgetSearchVariableCaseValue            string `json:"CAM_WIDGET_SEARCH_VARIABLE_CASE_VALUE"`
		CamWidgetSearchVariableCaseName             string `json:"CAM_WIDGET_SEARCH_VARIABLE_CASE_NAME"`
		CamWidgetSearchVariableCaseAttribute        string `json:"CAM_WIDGET_SEARCH_VARIABLE_CASE_ATTRIBUTE"`
		CamWidgetVariablesTableTooltipSetNull       string `json:"CAM_WIDGET_VARIABLES_TABLE_TOOLTIP_SET_NULL"`
		CamWidgetVariablesTablePlaceholderValue     string `json:"CAM_WIDGET_VARIABLES_TABLE_PLACEHOLDER_VALUE"`
		CamWidgetVariablesTableTooltipDefaultValue  string `json:"CAM_WIDGET_VARIABLES_TABLE_TOOLTIP_DEFAULT_VALUE"`
		CamWidgetVariablesTableTooltipUpload        string `json:"CAM_WIDGET_VARIABLES_TABLE_TOOLTIP_UPLOAD"`
		CamWidgetVariablesTableNull                 string `json:"CAM_WIDGET_VARIABLES_TABLE_NULL"`
		CamWidgetVariablesTableEditVariable         string `json:"CAM_WIDGET_VARIABLES_TABLE_EDIT_VARIABLE"`
		CamWidgetVariablesTableDeleteVariable       string `json:"CAM_WIDGET_VARIABLES_TABLE_DELETE_VARIABLE"`
		CamWidgetVariablesTableSaveVariable         string `json:"CAM_WIDGET_VARIABLES_TABLE_SAVE_VARIABLE"`
		CamWidgetVariablesTableRevertVariable       string `json:"CAM_WIDGET_VARIABLES_TABLE_REVERT_VARIABLE"`
		CamWidgetVariablesTableName                 string `json:"CAM_WIDGET_VARIABLES_TABLE_NAME"`
		CamWidgetVariablesTableType                 string `json:"CAM_WIDGET_VARIABLES_TABLE_TYPE"`
		CamWidgetVariablesTableValue                string `json:"CAM_WIDGET_VARIABLES_TABLE_VALUE"`
		CamWidgetVariablesTableDelete               string `json:"CAM_WIDGET_VARIABLES_TABLE_DELETE"`
		CamWidgetVariablesTableAbort                string `json:"CAM_WIDGET_VARIABLES_TABLE_ABORT"`
		CamWidgetVariablesTableDialogue             string `json:"CAM_WIDGET_VARIABLES_TABLE_DIALOGUE"`
		CamWidgetVariableDialogInspect              string `json:"CAM_WIDGET_VARIABLE_DIALOG_INSPECT"`
		CamWidgetVariableDialogLabelObjectType      string `json:"CAM_WIDGET_VARIABLE_DIALOG_LABEL_OBJECT_TYPE"`
		CamWidgetVariableDialogLabelSerialization   string `json:"CAM_WIDGET_VARIABLE_DIALOG_LABEL_SERIALIZATION"`
		CamWidgetVariableDialogLabelSerializedValue string `json:"CAM_WIDGET_VARIABLE_DIALOG_LABEL_SERIALIZED_VALUE"`
		CamWidgetVariableDialogBtnClose             string `json:"CAM_WIDGET_VARIABLE_DIALOG_BTN_CLOSE"`
		CamWidgetVariableDialogBtnChange            string `json:"CAM_WIDGET_VARIABLE_DIALOG_BTN_CHANGE"`
		CamWidgetVariableNull                       string `json:"CAM_WIDGET_VARIABLE_NULL"`
		CamWidgetVariableSetValueNull               string `json:"CAM_WIDGET_VARIABLE_SET_VALUE_NULL"`
		CamWidgetVariableValue                      string `json:"CAM_WIDGET_VARIABLE_VALUE"`
		CamWidgetVariableUndefined                  string `json:"CAM_WIDGET_VARIABLE_UNDEFINED"`
		CamWidgetVariableReset                      string `json:"CAM_WIDGET_VARIABLE_RESET"`
		CamWidgetStringDialogTitle                  string `json:"CAM_WIDGET_STRING_DIALOG_TITLE"`
		CamWidgetStringDialogLabelValue             string `json:"CAM_WIDGET_STRING_DIALOG_LABEL_VALUE"`
		CamWidgetStringDialogLabelClose             string `json:"CAM_WIDGET_STRING_DIALOG_LABEL_CLOSE"`
		Links                                       string `json:"LINKS"`
		Profile                                     string `json:"PROFILE"`
		EditProfile                                 string `json:"EDIT_PROFILE"`
		ChangePassword                              string `json:"CHANGE_PASSWORD"`
		EditYourProfile                             string `json:"EDIT_YOUR_PROFILE"`
		FirstName                                   string `json:"FIRST_NAME"`
		LastName                                    string `json:"LAST_NAME"`
		Email                                       string `json:"EMAIL"`
		Update                                      string `json:"UPDATE"`
		ChangeYourPassword                          string `json:"CHANGE_YOUR_PASSWORD"`
		CurrentPassword                             string `json:"CURRENT_PASSWORD"`
		NewPassword                                 string `json:"NEW_PASSWORD"`
		ConfirmNewPassword                          string `json:"CONFIRM_NEW_PASSWORD"`
		PasswordNotMatching                         string `json:"PASSWORD_NOT_MATCHING"`
		Groups                                      string `json:"GROUPS"`
		Username                                    string `json:"USERNAME"`
		ChangesSaved                                string `json:"CHANGES_SAVED"`
		ErrorWhileSaving                            string `json:"ERROR_WHILE_SAVING"`
		PasswordChanged                             string `json:"PASSWORD_CHANGED"`
		ErrorWhileChangingPassword                  string `json:"ERROR_WHILE_CHANGING_PASSWORD"`
		UsernameForAuthentication                   string `json:"USERNAME_FOR_AUTHENTICATION"`
		Applications                                string `json:"APPLICATIONS"`
		Cockpit                                     string `json:"COCKPIT"`
		Tasklist                                    string `json:"TASKLIST"`
		Admin                                       string `json:"ADMIN"`
		Documentation                               string `json:"DOCUMENTATION"`
		DocumentationDescription                    string `json:"DOCUMENTATION_DESCRIPTION"`
		PasswordPolicyTooltip                       string `json:"PASSWORD_POLICY_TOOLTIP"`
		PasswordPolicyTooltipPartial                string `json:"PASSWORD_POLICY_TOOLTIP_PARTIAL"`
		PasswordPolicyTooltipError                  string `json:"PASSWORD_POLICY_TOOLTIP_ERROR"`
		PasswordPolicyLength                        string `json:"PASSWORD_POLICY_LENGTH"`
		PasswordPolicyLowercase                     string `json:"PASSWORD_POLICY_LOWERCASE"`
		PasswordPolicyUppercase                     string `json:"PASSWORD_POLICY_UPPERCASE"`
		PasswordPolicyDigit                         string `json:"PASSWORD_POLICY_DIGIT"`
		PasswordPolicySpecial                       string `json:"PASSWORD_POLICY_SPECIAL"`
		PasswordPolicyUserData                      string `json:"PASSWORD_POLICY_USER_DATA"`
		FirstLoginInfo                              string `json:"FIRST_LOGIN_INFO"`
		FirstLoginHeading                           string `json:"FIRST_LOGIN_HEADING"`
		TelemetrySave                               string `json:"TELEMETRY_SAVE"`
		TelemetryCancel                             string `json:"TELEMETRY_CANCEL"`
		TelemetrySettings                           string `json:"TELEMETRY_SETTINGS"`
		TelemetryIntroduction                       string `json:"TELEMETRY_INTRODUCTION"`
		TelemetryEnable                             string `json:"TELEMETRY_ENABLE"`
		TelemetryDetails                            string `json:"TELEMETRY_DETAILS"`
		TelemetryPrivacy                            string `json:"TELEMETRY_PRIVACY"`
		TelemetryData                               string `json:"TELEMETRY_DATA"`
		TelemetryDocumentation                      string `json:"TELEMETRY_DOCUMENTATION"`
		TelemetryOrView                             string `json:"TELEMETRY_OR_VIEW"`
		TelemetryPrivacyPolicy                      string `json:"TELEMETRY_PRIVACY_POLICY"`
		TelemetrySuccessHeading                     string `json:"TELEMETRY_SUCCESS_HEADING"`
		TelemetrySuccess                            string `json:"TELEMETRY_SUCCESS"`
		TelemetryErrorStatus                        string `json:"TELEMETRY_ERROR_STATUS"`
		TelemetryErrorMessage                       string `json:"TELEMETRY_ERROR_MESSAGE"`
	}

func (w *WelcomeInterface) LoadFile(jsonFile string) {
	readfile := filepath.Join("lang", jsonFile)
	dat, err := ioutil.ReadFile(readfile)
	if err != nil {
		fmt.Println("No json file: ", err)
	}
	err = json.Unmarshal(dat, w)
	if err != nil {
		fmt.Println("Unmarshall error: ", err)
	}
}

func (w *WelcomeInterface) SaveWelcomeInterface(jsonFile string, lang string) {
	newFile := strings.Replace(jsonFile, "en", lang, -1)
	cwd, err := os.Getwd()
	dirLoc := filepath.Join(cwd, "output")
	saveFile := filepath.Join(dirLoc, newFile)
	err = os.Mkdir(dirLoc, 0755)
	if err != nil {
		if strings.Index(err.Error(), "file exists") == -1 {
			fmt.Println("No output folder: ", err)
		}
	}
	xfile, err := os.Create(saveFile)
	if err != nil {
		fmt.Println("Error creating file: ", err)
	}
	final, err := json.MarshalIndent(w, "", "    ")
	if err != nil {
		fmt.Println("Marshall error: ", err)
	}
	defer xfile.Close()
	xfile.Write(final)
	xfile.Close()
}

func (w *WelcomeInterface) TranslateWelcomeInterface(lang string) {
	mp := make(map[string]interface{})
	frLabels := make(map[string]interface{})
	reqBodyBytes := new(bytes.Buffer)
	json.NewEncoder(reqBodyBytes).Encode(w.Labels)
	err := json.Unmarshal(reqBodyBytes.Bytes(), &mp)
	if err != nil {
		fmt.Println("Unmarshall error: ", err)
	}
	trans := Translator{}
	for k, v := range mp {
		s, err := trans.TranslateTextWithModel(lang, fmt.Sprintf("%v", v), "nmt")
		if err != nil {
			fmt.Println("Translate Error: ", err)
		}
		frLabels[k] = s
	}
	labs := WelcomeLabels{}
	json.NewEncoder(reqBodyBytes).Encode(frLabels)
	bb, err := json.Marshal(frLabels)
	err = json.Unmarshal(bb, &labs)
	w.Labels = labs
}

