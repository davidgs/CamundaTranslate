package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

type AdminInterface struct {
	Labels AdminLabels `json:"labels,omitempty"`
}

type 	AdminLabels struct {
		AppVendor                                      string `json:"APP_VENDOR"`
		AdminApp                                       string `json:"ADMIN_APP"`
		TasklistApp                                    string `json:"TASKLIST_APP"`
		CockpitApp                                     string `json:"COCKPIT_APP"`
		ToggleNavigation                               string `json:"TOGGLE_NAVIGATION"`
		Submit                                         string `json:"SUBMIT"`
		Save                                           string `json:"SAVE"`
		Start                                          string `json:"START"`
		Abort                                          string `json:"ABORT"`
		Close                                          string `json:"CLOSE"`
		Cancel                                         string `json:"CANCEL"`
		Back                                           string `json:"BACK"`
		Delete                                         string `json:"DELETE"`
		Set                                            string `json:"SET"`
		Change                                         string `json:"CHANGE"`
		Task                                           string `json:"TASK"`
		Filter                                         string `json:"FILTER"`
		Tasks                                          string `json:"TASKS"`
		Filters                                        string `json:"FILTERS"`
		Loading                                        string `json:"LOADING"`
		Success                                        string `json:"SUCCESS"`
		Error                                          string `json:"ERROR"`
		RequiredField                                  string `json:"REQUIRED_FIELD"`
		RequiredIntegerField                           string `json:"REQUIRED_INTEGER_FIELD"`
		RequiredHexColorField                          string `json:"REQUIRED_HEX_COLOR_FIELD"`
		ProcessDefinition                              string `json:"PROCESS_DEFINITION"`
		ProcessInstance                                string `json:"PROCESS_INSTANCE"`
		ManageAccount                                  string `json:"MANAGE_ACCOUNT"`
		Asc                                            string `json:"ASC"`
		Desc                                           string `json:"DESC"`
		Boolean                                        string `json:"BOOLEAN"`
		Short                                          string `json:"SHORT"`
		Integer                                        string `json:"INTEGER"`
		Long                                           string `json:"LONG"`
		Double                                         string `json:"DOUBLE"`
		Date                                           string `json:"DATE"`
		String                                         string `json:"STRING"`
		AuthLogoutSuccessful                           string `json:"AUTH_LOGOUT_SUCCESSFUL"`
		AuthLogoutThanks                               string `json:"AUTH_LOGOUT_THANKS"`
		AuthDayContextMorning                          string `json:"AUTH_DAY_CONTEXT_MORNING"`
		AuthDayContextDay                              string `json:"AUTH_DAY_CONTEXT_DAY"`
		AuthDayContextAfternoon                        string `json:"AUTH_DAY_CONTEXT_AFTERNOON"`
		AuthDayContextEvening                          string `json:"AUTH_DAY_CONTEXT_EVENING"`
		AuthDayContextNight                            string `json:"AUTH_DAY_CONTEXT_NIGHT"`
		AuthDayContextWeekend                          string `json:"AUTH_DAY_CONTEXT_WEEKEND"`
		AuthFailedToDisplayResource                    string `json:"AUTH_FAILED_TO_DISPLAY_RESOURCE"`
		AuthAuthenticationFailed                       string `json:"AUTH_AUTHENTICATION_FAILED"`
		PageLoginErrorMsg                              string `json:"PAGE_LOGIN_ERROR_MSG"`
		PageLoginUsername                              string `json:"PAGE_LOGIN_USERNAME"`
		PageLoginPassword                              string `json:"PAGE_LOGIN_PASSWORD"`
		PageLoginPleaseSignIn                          string `json:"PAGE_LOGIN_PLEASE_SIGN_IN"`
		PageLoginSignInAction                          string `json:"PAGE_LOGIN_SIGN_IN_ACTION"`
		PageLoginFailed                                string `json:"PAGE_LOGIN_FAILED"`
		SignOutAction                                  string `json:"SIGN_OUT_ACTION"`
		MyProfile                                      string `json:"MY_PROFILE"`
		DirectiveEngineSelectTooltip                   string `json:"DIRECTIVE_ENGINE_SELECT_TOOLTIP"`
		DirectiveEngineSelectStatusNotFound            string `json:"DIRECTIVE_ENGINE_SELECT_STATUS_NOT_FOUND"`
		DirectiveEngineSelectMessageNotFound           string `json:"DIRECTIVE_ENGINE_SELECT_MESSAGE_NOT_FOUND"`
		DirectiveInplaceTextfieldErrorMsg              string `json:"DIRECTIVE_INPLACE_TEXTFIELD_ERROR_MSG"`
		PagesStatusServerError                         string `json:"PAGES_STATUS_SERVER_ERROR"`
		PagesMsgServerError                            string `json:"PAGES_MSG_SERVER_ERROR"`
		PagesStatusRequestTimeout                      string `json:"PAGES_STATUS_REQUEST_TIMEOUT"`
		PagesMsgRequestTimeout                         string `json:"PAGES_MSG_REQUEST_TIMEOUT"`
		PagesStatusAccessDenied                        string `json:"PAGES_STATUS_ACCESS_DENIED"`
		PagesMsgAccessDenied                           string `json:"PAGES_MSG_ACCESS_DENIED"`
		PagesMsgAccessDeniedResourceID                 string `json:"PAGES_MSG_ACCESS_DENIED_RESOURCE_ID"`
		PagesMsgActionDenied                           string `json:"PAGES_MSG_ACTION_DENIED"`
		PagesStatusNotFound                            string `json:"PAGES_STATUS_NOT_FOUND"`
		PagesMsgNotFound                               string `json:"PAGES_MSG_NOT_FOUND"`
		PagesStatusCommunicationError                  string `json:"PAGES_STATUS_COMMUNICATION_ERROR"`
		PagesMsgCommunicationError                     string `json:"PAGES_MSG_COMMUNICATION_ERROR"`
		PagesMsgEngineNotExists                        string `json:"PAGES_MSG_ENGINE_NOT_EXISTS"`
		ServicesResourceResolverIDNotFound             string `json:"SERVICES_RESOURCE_RESOLVER_ID_NOT_FOUND"`
		ServicesResourceResolverAuthFailed             string `json:"SERVICES_RESOURCE_RESOLVER_AUTH_FAILED"`
		ServicesResourceResolverReceivedStatus         string `json:"SERVICES_RESOURCE_RESOLVER_RECEIVED_STATUS"`
		ServicesResourceResolverDisplayFailed          string `json:"SERVICES_RESOURCE_RESOLVER_DISPLAY_FAILED"`
		ServicesResourceResolverNoPromise              string `json:"SERVICES_RESOURCE_RESOLVER_NO_PROMISE"`
		CamWidgetBpmnViewerError                       string `json:"CAM_WIDGET_BPMN_VIEWER_ERROR"`
		CamWidgetBpmnViewerLoading                     string `json:"CAM_WIDGET_BPMN_VIEWER_LOADING"`
		CamWidgetCopy                                  string `json:"CAM_WIDGET_COPY"`
		CamWidgetCopyLink                              string `json:"CAM_WIDGET_COPY_LINK"`
		CamWidgetCmmnViewerError                       string `json:"CAM_WIDGET_CMMN_VIEWER_ERROR"`
		CamWidgetCmmnViewerLoading                     string `json:"CAM_WIDGET_CMMN_VIEWER_LOADING"`
		CamWidgetDmnViewerError                        string `json:"CAM_WIDGET_DMN_VIEWER_ERROR"`
		CamWidgetDmnViewerLoading                      string `json:"CAM_WIDGET_DMN_VIEWER_LOADING"`
		CamWidgetFooterTimezone                        string `json:"CAM_WIDGET_FOOTER_TIMEZONE"`
		CamWidgetFooterPoweredBy                       string `json:"CAM_WIDGET_FOOTER_POWERED_BY"`
		CamWidgetHeaderToggleNavigation                string `json:"CAM_WIDGET_HEADER_TOGGLE_NAVIGATION"`
		CamWidgetHeaderMyProfile                       string `json:"CAM_WIDGET_HEADER_MY_PROFILE"`
		CamWidgetHeaderSignOut                         string `json:"CAM_WIDGET_HEADER_SIGN_OUT"`
		CamWidgetHeaderSmallScreenWarning              string `json:"CAM_WIDGET_HEADER_SMALL_SCREEN_WARNING"`
		CamWidgetSearchMatchType                       string `json:"CAM_WIDGET_SEARCH_MATCH_TYPE"`
		CamWidgetSearchMatchTypeAny                    string `json:"CAM_WIDGET_SEARCH_MATCH_TYPE_ANY"`
		CamWidgetSearchMatchTypeAll                    string `json:"CAM_WIDGET_SEARCH_MATCH_TYPE_ALL"`
		CamWidgetSearchTotalNumberResults              string `json:"CAM_WIDGET_SEARCH_TOTAL_NUMBER_RESULTS"`
		CamWidgetSearchVariableCaseValue               string `json:"CAM_WIDGET_SEARCH_VARIABLE_CASE_VALUE"`
		CamWidgetSearchVariableCaseName                string `json:"CAM_WIDGET_SEARCH_VARIABLE_CASE_NAME"`
		CamWidgetSearchVariableCaseAttribute           string `json:"CAM_WIDGET_SEARCH_VARIABLE_CASE_ATTRIBUTE"`
		CamWidgetVariablesTableTooltipSetNull          string `json:"CAM_WIDGET_VARIABLES_TABLE_TOOLTIP_SET_NULL"`
		CamWidgetVariablesTablePlaceholderValue        string `json:"CAM_WIDGET_VARIABLES_TABLE_PLACEHOLDER_VALUE"`
		CamWidgetVariablesTableTooltipDefaultValue     string `json:"CAM_WIDGET_VARIABLES_TABLE_TOOLTIP_DEFAULT_VALUE"`
		CamWidgetVariablesTableTooltipUpload           string `json:"CAM_WIDGET_VARIABLES_TABLE_TOOLTIP_UPLOAD"`
		CamWidgetVariablesTableNull                    string `json:"CAM_WIDGET_VARIABLES_TABLE_NULL"`
		CamWidgetVariablesTableEditVariable            string `json:"CAM_WIDGET_VARIABLES_TABLE_EDIT_VARIABLE"`
		CamWidgetVariablesTableDeleteVariable          string `json:"CAM_WIDGET_VARIABLES_TABLE_DELETE_VARIABLE"`
		CamWidgetVariablesTableSaveVariable            string `json:"CAM_WIDGET_VARIABLES_TABLE_SAVE_VARIABLE"`
		CamWidgetVariablesTableRevertVariable          string `json:"CAM_WIDGET_VARIABLES_TABLE_REVERT_VARIABLE"`
		CamWidgetVariablesTableName                    string `json:"CAM_WIDGET_VARIABLES_TABLE_NAME"`
		CamWidgetVariablesTableType                    string `json:"CAM_WIDGET_VARIABLES_TABLE_TYPE"`
		CamWidgetVariablesTableValue                   string `json:"CAM_WIDGET_VARIABLES_TABLE_VALUE"`
		CamWidgetVariablesTableDelete                  string `json:"CAM_WIDGET_VARIABLES_TABLE_DELETE"`
		CamWidgetVariablesTableAbort                   string `json:"CAM_WIDGET_VARIABLES_TABLE_ABORT"`
		CamWidgetVariablesTableDialogue                string `json:"CAM_WIDGET_VARIABLES_TABLE_DIALOGUE"`
		CamWidgetVariableDialogInspect                 string `json:"CAM_WIDGET_VARIABLE_DIALOG_INSPECT"`
		CamWidgetVariableDialogLabelObjectType         string `json:"CAM_WIDGET_VARIABLE_DIALOG_LABEL_OBJECT_TYPE"`
		CamWidgetVariableDialogLabelSerialization      string `json:"CAM_WIDGET_VARIABLE_DIALOG_LABEL_SERIALIZATION"`
		CamWidgetVariableDialogLabelSerializedValue    string `json:"CAM_WIDGET_VARIABLE_DIALOG_LABEL_SERIALIZED_VALUE"`
		CamWidgetVariableDialogBtnClose                string `json:"CAM_WIDGET_VARIABLE_DIALOG_BTN_CLOSE"`
		CamWidgetVariableDialogBtnChange               string `json:"CAM_WIDGET_VARIABLE_DIALOG_BTN_CHANGE"`
		CamWidgetVariableNull                          string `json:"CAM_WIDGET_VARIABLE_NULL"`
		CamWidgetVariableSetValueNull                  string `json:"CAM_WIDGET_VARIABLE_SET_VALUE_NULL"`
		CamWidgetVariableValue                         string `json:"CAM_WIDGET_VARIABLE_VALUE"`
		CamWidgetVariableUndefined                     string `json:"CAM_WIDGET_VARIABLE_UNDEFINED"`
		CamWidgetVariableReset                         string `json:"CAM_WIDGET_VARIABLE_RESET"`
		CamWidgetStringDialogTitle                     string `json:"CAM_WIDGET_STRING_DIALOG_TITLE"`
		CamWidgetStringDialogLabelValue                string `json:"CAM_WIDGET_STRING_DIALOG_LABEL_VALUE"`
		CamWidgetStringDialogLabelClose                string `json:"CAM_WIDGET_STRING_DIALOG_LABEL_CLOSE"`
		AuthorizationUpdate                            string `json:"AUTHORIZATION_UPDATE"`
		AuthorizationCreate                            string `json:"AUTHORIZATION_CREATE"`
		AuthorizationGlobal                            string `json:"AUTHORIZATION_GLOBAL"`
		AuthorizationAllow                             string `json:"AUTHORIZATION_ALLOW"`
		AuthorizationDeny                              string `json:"AUTHORIZATION_DENY"`
		AuthorizationAuthorizations                    string `json:"AUTHORIZATION_AUTHORIZATIONS"`
		AuthorizationCreateNewAuthorization            string `json:"AUTHORIZATION_CREATE_NEW_AUTHORIZATION"`
		AuthorizationType                              string `json:"AUTHORIZATION_TYPE"`
		AuthorizationUserGroup                         string `json:"AUTHORIZATION_USER_GROUP"`
		AuthorizationUserGroupID                       string `json:"AUTHORIZATION_USER_GROUP_ID"`
		AuthorizationPermissions                       string `json:"AUTHORIZATION_PERMISSIONS"`
		AuthorizationResourceID                        string `json:"AUTHORIZATION_RESOURCE_ID"`
		AuthorizationAction                            string `json:"AUTHORIZATION_ACTION"`
		AuthorizationCockpitOrTasklist                 string `json:"AUTHORIZATION_COCKPIT_OR_TASKLIST"`
		AuthorizationEdit                              string `json:"AUTHORIZATION_EDIT"`
		AuthorizationDelete                            string `json:"AUTHORIZATION_DELETE"`
		AuthorizationApplication                       string `json:"AUTHORIZATION_APPLICATION"`
		AuthorizationUser                              string `json:"AUTHORIZATION_USER"`
		AuthorizationGroup                             string `json:"AUTHORIZATION_GROUP"`
		AuthorizationGroupMembership                   string `json:"AUTHORIZATION_GROUP_MEMBERSHIP"`
		AuthorizationAuthorization                     string `json:"AUTHORIZATION_AUTHORIZATION"`
		AuthorizationFilter                            string `json:"AUTHORIZATION_FILTER"`
		AuthorizationProcessDefinition                 string `json:"AUTHORIZATION_PROCESS_DEFINITION"`
		AuthorizationTask                              string `json:"AUTHORIZATION_TASK"`
		AuthorizationProcessInstance                   string `json:"AUTHORIZATION_PROCESS_INSTANCE"`
		AuthorizationDeployment                        string `json:"AUTHORIZATION_DEPLOYMENT"`
		AuthorizationDecisionDefinition                string `json:"AUTHORIZATION_DECISION_DEFINITION"`
		AuthorizationTenant                            string `json:"AUTHORIZATION_TENANT"`
		AuthorizationTenantMembership                  string `json:"AUTHORIZATION_TENANT_MEMBERSHIP"`
		AuthorizationBatch                             string `json:"AUTHORIZATION_BATCH"`
		AuthorizationDecisionRequirementsDefinition    string `json:"AUTHORIZATION_DECISION_REQUIREMENTS_DEFINITION"`
		AuthorizationReport                            string `json:"AUTHORIZATION_REPORT"`
		AuthorizationDashboard                         string `json:"AUTHORIZATION_DASHBOARD"`
		AuthorizationOptimize                          string `json:"AUTHORIZATION_OPTIMIZE"`
		AuthorizationConfirmDeleteTitle                string `json:"AUTHORIZATION_CONFIRM_DELETE_TITLE"`
		AuthorizationConfirmDeleteMessage              string `json:"AUTHORIZATION_CONFIRM_DELETE_MESSAGE"`
		AuthorizationResource                          string `json:"AUTHORIZATION_RESOURCE"`
		AuthorizationDeleteMessage                     string `json:"AUTHORIZATION_DELETE_MESSAGE"`
		AuthorizationDeleteClose                       string `json:"AUTHORIZATION_DELETE_CLOSE"`
		AuthorizationDeleteOk                          string `json:"AUTHORIZATION_DELETE_OK"`
		AuthorizationDeleteDelete                      string `json:"AUTHORIZATION_DELETE_DELETE"`
		AuthorizationManageAuthorizations              string `json:"AUTHORIZATION_MANAGE_AUTHORIZATIONS"`
		AuthorizationOperationLog                      string `json:"AUTHORIZATION_OPERATION_LOG"`
		AuthorizationHistoricTask                      string `json:"AUTHORIZATION_HISTORIC_TASK"`
		AuthorizationHistoricProcessInstance           string `json:"AUTHORIZATION_HISTORIC_PROCESS_INSTANCE"`
		TelemetrySave                                  string `json:"TELEMETRY_SAVE"`
		TelemetryCancel                                string `json:"TELEMETRY_CANCEL"`
		TelemetrySettings                              string `json:"TELEMETRY_SETTINGS"`
		TelemetryIntroduction                          string `json:"TELEMETRY_INTRODUCTION"`
		TelemetryEnable                                string `json:"TELEMETRY_ENABLE"`
		TelemetryDetails                               string `json:"TELEMETRY_DETAILS"`
		TelemetryPrivacy                               string `json:"TELEMETRY_PRIVACY"`
		TelemetryData                                  string `json:"TELEMETRY_DATA"`
		TelemetryDocumentation                         string `json:"TELEMETRY_DOCUMENTATION"`
		TelemetryOrView                                string `json:"TELEMETRY_OR_VIEW"`
		TelemetryPrivacyPolicy                         string `json:"TELEMETRY_PRIVACY_POLICY"`
		TelemetrySuccessHeading                        string `json:"TELEMETRY_SUCCESS_HEADING"`
		TelemetrySuccess                               string `json:"TELEMETRY_SUCCESS"`
		TelemetryErrorStatus                           string `json:"TELEMETRY_ERROR_STATUS"`
		TelemetryErrorMessage                          string `json:"TELEMETRY_ERROR_MESSAGE"`
		GroupMembershipGroupSelectGroups               string `json:"GROUP_MEMBERSHIP_GROUP_SELECT_GROUPS"`
		GroupMembershipGroupGroupID                    string `json:"GROUP_MEMBERSHIP_GROUP_GROUP_ID"`
		GroupMembershipGroupGroupName                  string `json:"GROUP_MEMBERSHIP_GROUP_GROUP_NAME"`
		GroupMembershipGroupGroupType                  string `json:"GROUP_MEMBERSHIP_GROUP_GROUP_TYPE"`
		GroupMembershipGroupMessageAvailableGroups     string `json:"GROUP_MEMBERSHIP_GROUP_MESSAGE_AVAILABLE_GROUPS"`
		GroupMembershipGroupMessageAvailableGroupsPage string `json:"GROUP_MEMBERSHIP_GROUP_MESSAGE_AVAILABLE_GROUPS_PAGE"`
		GroupMembershipGroupMessageHere                string `json:"GROUP_MEMBERSHIP_GROUP_MESSAGE_HERE"`
		GroupMembershipGroupMessageSuccess             string `json:"GROUP_MEMBERSHIP_GROUP_MESSAGE_SUCCESS"`
		GroupMembershipGroupClose                      string `json:"GROUP_MEMBERSHIP_GROUP_CLOSE"`
		GroupMembershipGroupOk                         string `json:"GROUP_MEMBERSHIP_GROUP_OK"`
		GroupMembershipGroupAddGroups                  string `json:"GROUP_MEMBERSHIP_GROUP_ADD_GROUPS"`
		GroupMembershipTenantSelectTenants             string `json:"GROUP_MEMBERSHIP_TENANT_SELECT_TENANTS"`
		GroupMembershipTenantTenantID                  string `json:"GROUP_MEMBERSHIP_TENANT_TENANT_ID"`
		GroupMembershipTenantTenantName                string `json:"GROUP_MEMBERSHIP_TENANT_TENANT_NAME"`
		GroupMembershipTenantMessageAvailableTenants   string `json:"GROUP_MEMBERSHIP_TENANT_MESSAGE_AVAILABLE_TENANTS"`
		GroupMembershipTenantMessageHere               string `json:"GROUP_MEMBERSHIP_TENANT_MESSAGE_HERE"`
		GroupMembershipTenantMessageSuccess            string `json:"GROUP_MEMBERSHIP_TENANT_MESSAGE_SUCCESS"`
		GroupMembershipTenantClose                     string `json:"GROUP_MEMBERSHIP_TENANT_CLOSE"`
		GroupMembershipTenantOk                        string `json:"GROUP_MEMBERSHIP_TENANT_OK"`
		GroupMembershipTenantAdd                       string `json:"GROUP_MEMBERSHIP_TENANT_ADD"`
		UserMembershipTenantSelectTenants              string `json:"USER_MEMBERSHIP_TENANT_SELECT_TENANTS"`
		UserMembershipTenantTenantID                   string `json:"USER_MEMBERSHIP_TENANT_TENANT_ID"`
		UserMembershipTenantTenantName                 string `json:"USER_MEMBERSHIP_TENANT_TENANT_NAME"`
		UserMembershipTenantMessageAvailableTenants    string `json:"USER_MEMBERSHIP_TENANT_MESSAGE_AVAILABLE_TENANTS"`
		UserMembershipTenantMessageHere                string `json:"USER_MEMBERSHIP_TENANT_MESSAGE_HERE"`
		UserMembershipTenantMessageSuccess             string `json:"USER_MEMBERSHIP_TENANT_MESSAGE_SUCCESS"`
		UserMembershipTenantClose                      string `json:"USER_MEMBERSHIP_TENANT_CLOSE"`
		UserMembershipTenantOk                         string `json:"USER_MEMBERSHIP_TENANT_OK"`
		UserMembershipTenantAdd                        string `json:"USER_MEMBERSHIP_TENANT_ADD"`
		ExecutionMetrics                               string `json:"EXECUTION_METRICS"`
		ExecutionMetricsInfo1                          string `json:"EXECUTION_METRICS_INFO_1"`
		ExecutionMetricsFni                            string `json:"EXECUTION_METRICS_FNI"`
		ExecutionMetricsAnd                            string `json:"EXECUTION_METRICS_AND"`
		ExecutionMetricsEde                            string `json:"EXECUTION_METRICS_EDE"`
		ExecutionMetricsRpi                            string `json:"EXECUTION_METRICS_RPI"`
		ExecutionMetricsDi                             string `json:"EXECUTION_METRICS_DI"`
		ExecutionMetricsTw                             string `json:"EXECUTION_METRICS_TW"`
		ExecutionMetricsTaskWorkersShow                string `json:"EXECUTION_METRICS_TASK_WORKERS_SHOW"`
		ExecutionMetricsInfo2                          string `json:"EXECUTION_METRICS_INFO_2"`
		ExecutionMetricsInfo3                          string `json:"EXECUTION_METRICS_INFO_3"`
		ExecutionMetricsStartDate                      string `json:"EXECUTION_METRICS_START_DATE"`
		ExecutionMetricsEndDate                        string `json:"EXECUTION_METRICS_END_DATE"`
		ExecutionMetricsResults                        string `json:"EXECUTION_METRICS_RESULTS"`
		ExecutionMetricsRefresh                        string `json:"EXECUTION_METRICS_REFRESH"`
		AnErrorOccurred                                string `json:"AN_ERROR_OCCURRED"`
		GenericConfirmationTitle                       string `json:"GENERIC_CONFIRMATION_TITLE"`
		GenericConfirmationCancel                      string `json:"GENERIC_CONFIRMATION_CANCEL"`
		GenericConfirmationProceed                     string `json:"GENERIC_CONFIRMATION_PROCEED"`
		GroupCreateNewGroup                            string `json:"GROUP_CREATE_NEW_GROUP"`
		GroupCreateGroupID                             string `json:"GROUP_CREATE_GROUP_ID"`
		GroupCreateIDRequired                          string `json:"GROUP_CREATE_ID_REQUIRED"`
		GroupCreateGroupName                           string `json:"GROUP_CREATE_GROUP_NAME"`
		GroupCreateNameRequired                        string `json:"GROUP_CREATE_NAME_REQUIRED"`
		GroupGroupType                                 string `json:"GROUP_GROUP_TYPE"`
		GroupCreateCancel                              string `json:"GROUP_CREATE_CANCEL"`
		GroupCreateLabelGroup                          string `json:"GROUP_CREATE_LABEL_GROUP"`
		GroupCreateLabelNewGroup                       string `json:"GROUP_CREATE_LABEL_NEW_GROUP"`
		GroupCreateMessageSuccess                      string `json:"GROUP_CREATE_MESSAGE_SUCCESS"`
		GroupCreateMessageError                        string `json:"GROUP_CREATE_MESSAGE_ERROR"`
		GroupEditInformation                           string `json:"GROUP_EDIT_INFORMATION"`
		GroupEditTenants                               string `json:"GROUP_EDIT_TENANTS"`
		GroupEditTenantsLink                           string `json:"GROUP_EDIT_TENANTS_LINK"`
		GroupEditUsers                                 string `json:"GROUP_EDIT_USERS"`
		GroupEditEditGroup                             string `json:"GROUP_EDIT_EDIT_GROUP"`
		GroupEditAddToTenant                           string `json:"GROUP_EDIT_ADD_TO_TENANT"`
		GroupEditGroupDetails                          string `json:"GROUP_EDIT_GROUP_DETAILS"`
		GroupEditGroupName                             string `json:"GROUP_EDIT_GROUP_NAME"`
		GroupEditNameRequired                          string `json:"GROUP_EDIT_NAME_REQUIRED"`
		GroupEditGroupType                             string `json:"GROUP_EDIT_GROUP_TYPE"`
		GroupEditUpdateGroup                           string `json:"GROUP_EDIT_UPDATE_GROUP"`
		GroupEditDeleteGroup                           string `json:"GROUP_EDIT_DELETE_GROUP"`
		GroupEditWarning                               string `json:"GROUP_EDIT_WARNING"`
		GroupEditWarningDelete                         string `json:"GROUP_EDIT_WARNING_DELETE"`
		GroupEditMessageGroup                          string `json:"GROUP_EDIT_MESSAGE_GROUP"`
		GroupEditTenantID                              string `json:"GROUP_EDIT_TENANT_ID"`
		GroupEditTenantName                            string `json:"GROUP_EDIT_TENANT_NAME"`
		GroupEditAction                                string `json:"GROUP_EDIT_ACTION"`
		GroupEditRemove                                string `json:"GROUP_EDIT_REMOVE"`
		GroupEditGroupUsers                            string `json:"GROUP_EDIT_GROUP_USERS"`
		GroupEditMessageUsers                          string `json:"GROUP_EDIT_MESSAGE_USERS"`
		GroupEditName                                  string `json:"GROUP_EDIT_NAME"`
		GroupEditUsername                              string `json:"GROUP_EDIT_USERNAME"`
		GroupEditGroup                                 string `json:"GROUP_EDIT_GROUP"`
		GroupEditGroups                                string `json:"GROUP_EDIT_GROUPS"`
		GroupEditRemovedFromTenant                     string `json:"GROUP_EDIT_REMOVED_FROM_TENANT"`
		GroupEditUpdateSuccess                         string `json:"GROUP_EDIT_UPDATE_SUCCESS"`
		GroupEditUpdateFailed                          string `json:"GROUP_EDIT_UPDATE_FAILED"`
		GroupEditDeleteConfirm                         string `json:"GROUP_EDIT_DELETE_CONFIRM"`
		GroupEditDeleteSuccess                         string `json:"GROUP_EDIT_DELETE_SUCCESS"`
		GroupEditDeleteFailed                          string `json:"GROUP_EDIT_DELETE_FAILED"`
		GroupMembershipCreateLoadFailed                string `json:"GROUP_MEMBERSHIP_CREATE_LOAD_FAILED"`
		GroupMembershipCreateCreateFailed              string `json:"GROUP_MEMBERSHIP_CREATE_CREATE_FAILED"`
		GroupsListGroups                               string `json:"GROUPS_LIST_GROUPS"`
		GroupsCreateGroup                              string `json:"GROUPS_CREATE_GROUP"`
		GroupsGroupID                                  string `json:"GROUPS_GROUP_ID"`
		GroupsGroupName                                string `json:"GROUPS_GROUP_NAME"`
		GroupsGroupType                                string `json:"GROUPS_GROUP_TYPE"`
		GroupsAction                                   string `json:"GROUPS_ACTION"`
		GroupsActionEdit                               string `json:"GROUPS_ACTION_EDIT"`
		GroupsGroups                                   string `json:"GROUPS_GROUPS"`
		GroupsCreateNew                                string `json:"GROUPS_CREATE_NEW"`
		GroupsListOfGroups                             string `json:"GROUPS_LIST_OF_GROUPS"`
		GroupsNoGroups                                 string `json:"GROUPS_NO_GROUPS"`
		GroupsGroup                                    string `json:"GROUPS_GROUP"`
		SetupSetup                                     string `json:"SETUP_SETUP"`
		SetupUserCreated                               string `json:"SETUP_USER_CREATED"`
		SetupUserInitial                               string `json:"SETUP_USER_INITIAL"`
		SetupUserRedirect1                             string `json:"SETUP_USER_REDIRECT_1"`
		SetupUserRedirect2                             string `json:"SETUP_USER_REDIRECT_2"`
		SetupUserRedirect3                             string `json:"SETUP_USER_REDIRECT_3"`
		SetupUserAccount                               string `json:"SETUP_USER_ACCOUNT"`
		SetupUserID                                    string `json:"SETUP_USER_ID"`
		SetupUserIDRequired                            string `json:"SETUP_USER_ID_REQUIRED"`
		SetupUserPassword                              string `json:"SETUP_USER_PASSWORD"`
		SetupUserPasswordLength                        string `json:"SETUP_USER_PASSWORD_LENGTH"`
		SetupUserPasswordRequired                      string `json:"SETUP_USER_PASSWORD_REQUIRED"`
		SetupUserPasswordRepeat                        string `json:"SETUP_USER_PASSWORD_REPEAT"`
		SetupUserPasswordMustMatch                     string `json:"SETUP_USER_PASSWORD_MUST_MATCH"`
		SetupUserProfile                               string `json:"SETUP_USER_PROFILE"`
		SetupUserFirstname                             string `json:"SETUP_USER_FIRSTNAME"`
		SetupUserFirstnameRequired                     string `json:"SETUP_USER_FIRSTNAME_REQUIRED"`
		SetupUserLastname                              string `json:"SETUP_USER_LASTNAME"`
		SetupUserLastnameRequired                      string `json:"SETUP_USER_LASTNAME_REQUIRED"`
		SetupUserEmail                                 string `json:"SETUP_USER_EMAIL"`
		SetupUserEmailNotValid                         string `json:"SETUP_USER_EMAIL_NOT_VALID"`
		SetupUserNewUser                               string `json:"SETUP_USER_NEW_USER"`
		SetupUserMessage1                              string `json:"SETUP_USER_MESSAGE_1"`
		SetupUserMessage2                              string `json:"SETUP_USER_MESSAGE_2"`
		SetupUserMessage3                              string `json:"SETUP_USER_MESSAGE_3"`
		SetupUserMessage4                              string `json:"SETUP_USER_MESSAGE_4"`
		SetupUserMessage5                              string `json:"SETUP_USER_MESSAGE_5"`
		SetupCouldNotCreateUser                        string `json:"SETUP_COULD_NOT_CREATE_USER"`
		SystemSystemSettings                           string `json:"SYSTEM_SYSTEM_SETTINGS"`
		SystemGeneralSettings                          string `json:"SYSTEM_GENERAL_SETTINGS"`
		SystemProcessEngine                            string `json:"SYSTEM_PROCESS_ENGINE"`
		SystemProcessRunning                           string `json:"SYSTEM_PROCESS_RUNNING"`
		SystemGeneral                                  string `json:"SYSTEM_GENERAL"`
		UsersCreateNew                                 string `json:"USERS_CREATE_NEW"`
		UsersListUsers                                 string `json:"USERS_LIST_USERS"`
		UsersMyProfile                                 string `json:"USERS_MY_PROFILE"`
		UsersListOfUsers                               string `json:"USERS_LIST_OF_USERS"`
		UsersAddUser                                   string `json:"USERS_ADD_USER"`
		UsersNotFound                                  string `json:"USERS_NOT_FOUND"`
		UsersName                                      string `json:"USERS_NAME"`
		UsersUsername                                  string `json:"USERS_USERNAME"`
		UsersAction                                    string `json:"USERS_ACTION"`
		UsersEdit                                      string `json:"USERS_EDIT"`
		UsersEditUser                                  string `json:"USERS_EDIT_USER"`
		UsersEditSuccessMsn                            string `json:"USERS_EDIT_SUCCESS_MSN"`
		UsersEditFailed                                string `json:"USERS_EDIT_FAILED"`
		UsersPasswordChanged                           string `json:"USERS_PASSWORD_CHANGED"`
		UsersOldPasswordNotValid                       string `json:"USERS_OLD_PASSWORD_NOT_VALID"`
		UsersPasswordNotValid                          string `json:"USERS_PASSWORD_NOT_VALID"`
		UsersPasswordCouldNotChange                    string `json:"USERS_PASSWORD_COULD_NOT_CHANGE"`
		UsersUserDeleteConfirm                         string `json:"USERS_USER_DELETE_CONFIRM"`
		UsersUserDeleteSuccess                         string `json:"USERS_USER_DELETE_SUCCESS"`
		UsersUserDeleteFromGroup                       string `json:"USERS_USER_DELETE_FROM_GROUP"`
		UsersUserDeleteFromTenant                      string `json:"USERS_USER_DELETE_FROM_TENANT"`
		UsersUserUnlockSuccess                         string `json:"USERS_USER_UNLOCK_SUCCESS"`
		UsersUsers                                     string `json:"USERS_USERS"`
		UsersProfile                                   string `json:"USERS_PROFILE"`
		UsersAccount                                   string `json:"USERS_ACCOUNT"`
		UsersGroups                                    string `json:"USERS_GROUPS"`
		UsersFirstnameLastnameGroups                   string `json:"USERS_FIRSTNAME_LASTNAME_GROUPS"`
		UsersTenants                                   string `json:"USERS_TENANTS"`
		UsersFirstnameLastnameTenants                  string `json:"USERS_FIRSTNAME_LASTNAME_TENANTS"`
		UsersAddToAGroup                               string `json:"USERS_ADD_TO_A_GROUP"`
		UsersAddToATenant                              string `json:"USERS_ADD_TO_A_TENANT"`
		UsersLoadingProfile                            string `json:"USERS_LOADING_PROFILE"`
		UsersID                                        string `json:"USERS_ID"`
		UsersFirstname                                 string `json:"USERS_FIRSTNAME"`
		UsersFirstnameRequired                         string `json:"USERS_FIRSTNAME_REQUIRED"`
		UsersLastname                                  string `json:"USERS_LASTNAME"`
		UsersLastnameRequired                          string `json:"USERS_LASTNAME_REQUIRED"`
		UsersEmail                                     string `json:"USERS_EMAIL"`
		UsersEmailInvalid                              string `json:"USERS_EMAIL_INVALID"`
		UsersUpdateProfile                             string `json:"USERS_UPDATE_PROFILE"`
		UsersChangePassword                            string `json:"USERS_CHANGE_PASSWORD"`
		UsersTypeNewPassword                           string `json:"USERS_TYPE_NEW_PASSWORD"`
		UsersOldPassword                               string `json:"USERS_OLD_PASSWORD"`
		UsersMyPassword                                string `json:"USERS_MY_PASSWORD"`
		UsersNewPassword                               string `json:"USERS_NEW_PASSWORD"`
		UsersPasswordInvalid                           string `json:"USERS_PASSWORD_INVALID"`
		UsersNewPasswordRepeat                         string `json:"USERS_NEW_PASSWORD_REPEAT"`
		UsersPasswordsNotEqual                         string `json:"USERS_PASSWORDS_NOT_EQUAL"`
		UsersDeleteUser                                string `json:"USERS_DELETE_USER"`
		UsersWarning                                   string `json:"USERS_WARNING"`
		UsersDeletingUndone                            string `json:"USERS_DELETING_UNDONE"`
		UsersNotice                                    string `json:"USERS_NOTICE"`
		UsersUnlockingUser                             string `json:"USERS_UNLOCKING_USER"`
		UsersUnlockUser                                string `json:"USERS_UNLOCK_USER"`
		UsersGroupID                                   string `json:"USERS_GROUP_ID"`
		UsersGroupName                                 string `json:"USERS_GROUP_NAME"`
		UsersGroupType                                 string `json:"USERS_GROUP_TYPE"`
		UsersRemove                                    string `json:"USERS_REMOVE"`
		UsersTenantID                                  string `json:"USERS_TENANT_ID"`
		UsersTenantName                                string `json:"USERS_TENANT_NAME"`
		UsersUser                                      string `json:"USERS_USER"`
		UsersNotMemberTentant                          string `json:"USERS_NOT_MEMBER_TENTANT"`
		UsersNotMemberGroup                            string `json:"USERS_NOT_MEMBER_GROUP"`
		UsersUserAccount                               string `json:"USERS_USER_ACCOUNT"`
		UsersUserID                                    string `json:"USERS_USER_ID"`
		UsersUserIDRequired                            string `json:"USERS_USER_ID_REQUIRED"`
		UsersPassword                                  string `json:"USERS_PASSWORD"`
		UsersPasswordRequired                          string `json:"USERS_PASSWORD_REQUIRED"`
		UsersPasswordRepeat                            string `json:"USERS_PASSWORD_REPEAT"`
		UsersUserProfile                               string `json:"USERS_USER_PROFILE"`
		UsersCreateNewUser                             string `json:"USERS_CREATE_NEW_USER"`
		UsersCancel                                    string `json:"USERS_CANCEL"`
		UsersNoSortingHeader                           string `json:"USERS_NO_SORTING_HEADER"`
		UsersNoSortingBody                             string `json:"USERS_NO_SORTING_BODY"`
		TenantsTenants                                 string `json:"TENANTS_TENANTS"`
		TenantsListTenants                             string `json:"TENANTS_LIST_TENANTS"`
		TenantsListOfTenants                           string `json:"TENANTS_LIST_OF_TENANTS"`
		TenantsCreateNew                               string `json:"TENANTS_CREATE_NEW"`
		TenantsCreateTenant                            string `json:"TENANTS_CREATE_TENANT"`
		TenantsNoTenants                               string `json:"TENANTS_NO_TENANTS"`
		TenantsAction                                  string `json:"TENANTS_ACTION"`
		TenantsEdit                                    string `json:"TENANTS_EDIT"`
		TenantsInformation                             string `json:"TENANTS_INFORMATION"`
		TenantsGroups                                  string `json:"TENANTS_GROUPS"`
		TenantsUsers                                   string `json:"TENANTS_USERS"`
		TenantsTenantName                              string `json:"TENANTS_TENANT_NAME"`
		TenantsNameRequired                            string `json:"TENANTS_NAME_REQUIRED"`
		TenantsUpdateTenant                            string `json:"TENANTS_UPDATE_TENANT"`
		TenantsDeleteTenant                            string `json:"TENANTS_DELETE_TENANT"`
		TenantsWarning                                 string `json:"TENANTS_WARNING"`
		TenantsDeletingTenantUndone                    string `json:"TENANTS_DELETING_TENANT_UNDONE"`
		TenantsTenantUsers                             string `json:"TENANTS_TENANT_USERS"`
		TenantsNoUsers                                 string `json:"TENANTS_NO_USERS"`
		TenantsName                                    string `json:"TENANTS_NAME"`
		TenantsUsername                                string `json:"TENANTS_USERNAME"`
		TenantsTenantGroups                            string `json:"TENANTS_TENANT_GROUPS"`
		TenantsNoGroups                                string `json:"TENANTS_NO_GROUPS"`
		TenantsTenantID                                string `json:"TENANTS_TENANT_ID"`
		TenantsIDRequired                              string `json:"TENANTS_ID_REQUIRED"`
		TenantsCancel                                  string `json:"TENANTS_CANCEL"`
		TenantsCreateTenantSuccess                     string `json:"TENANTS_CREATE_TENANT_SUCCESS"`
		TenantsCreateTenantFailed                      string `json:"TENANTS_CREATE_TENANT_FAILED"`
		TenantsTenant                                  string `json:"TENANTS_TENANT"`
		TenantsTenantTitle                             string `json:"TENANTS_TENANT_TITLE"`
		TenantsTenantUpdateSuccess                     string `json:"TENANTS_TENANT_UPDATE_SUCCESS"`
		TenantsTenantUpdateFailed                      string `json:"TENANTS_TENANT_UPDATE_FAILED"`
		TenantsTenantDeleteConfirm                     string `json:"TENANTS_TENANT_DELETE_CONFIRM"`
		TenantsTenantDeleteSuccess                     string `json:"TENANTS_TENANT_DELETE_SUCCESS"`
		TenantsTenantDeleteFailed                      string `json:"TENANTS_TENANT_DELETE_FAILED"`
		TenantsTenantLoadFailed                        string `json:"TENANTS_TENANT_LOAD_FAILED"`
		TenantsTenantCreateFailed                      string `json:"TENANTS_TENANT_CREATE_FAILED"`
		UsersCreateUser                                string `json:"USERS_CREATE_USER"`
		UsersCreate                                    string `json:"USERS_CREATE"`
		UsersCreateSuccess                             string `json:"USERS_CREATE_SUCCESS"`
		UsersCreateFailed                              string `json:"USERS_CREATE_FAILED"`
		SystemSystem                                   string `json:"SYSTEM_SYSTEM"`
		NotificationsStatusSuccess                     string `json:"NOTIFICATIONS_STATUS_SUCCESS"`
		NotificationsStatusFailed                      string `json:"NOTIFICATIONS_STATUS_FAILED"`
		NotificationsStatusError                       string `json:"NOTIFICATIONS_STATUS_ERROR"`
		NotificationsStatusPassword                    string `json:"NOTIFICATIONS_STATUS_PASSWORD"`
		DashboardDashboard                             string `json:"DASHBOARD_DASHBOARD"`
		SearchPlaceholder                              string `json:"SEARCH_PLACEHOLDER"`
		InvalidSearch                                  string `json:"INVALID_SEARCH"`
		DeleteSearch                                   string `json:"DELETE_SEARCH"`
		Type                                           string `json:"TYPE"`
		Property                                       string `json:"PROPERTY"`
		Operator                                       string `json:"OPERATOR"`
		Value                                          string `json:"VALUE"`
		PluginSearchGroupsID                           string `json:"PLUGIN_SEARCH_GROUPS_ID"`
		PluginSearchGroupsName                         string `json:"PLUGIN_SEARCH_GROUPS_NAME"`
		PluginSearchGroupsLike                         string `json:"PLUGIN_SEARCH_GROUPS_LIKE"`
		PluginSearchGroupsType                         string `json:"PLUGIN_SEARCH_GROUPS_TYPE"`
		PluginSearchGroupsUserID                       string `json:"PLUGIN_SEARCH_GROUPS_USER_ID"`
		PluginSearchGroupsTenantID                     string `json:"PLUGIN_SEARCH_GROUPS_TENANT_ID"`
		PluginSearchTenantsID                          string `json:"PLUGIN_SEARCH_TENANTS_ID"`
		PluginSearchTenantsName                        string `json:"PLUGIN_SEARCH_TENANTS_NAME"`
		PluginSearchTenantsLike                        string `json:"PLUGIN_SEARCH_TENANTS_LIKE"`
		PluginSearchTenantsUserID                      string `json:"PLUGIN_SEARCH_TENANTS_USER_ID"`
		PluginSearchTenantsGroupID                     string `json:"PLUGIN_SEARCH_TENANTS_GROUP_ID"`
		PluginSearchUserID                             string `json:"PLUGIN_SEARCH_USER_ID"`
		PluginSearchUserFirstname                      string `json:"PLUGIN_SEARCH_USER_FIRSTNAME"`
		PluginSearchUserLastname                       string `json:"PLUGIN_SEARCH_USER_LASTNAME"`
		PluginSearchUserLike                           string `json:"PLUGIN_SEARCH_USER_LIKE"`
		PluginSearchUserEmail                          string `json:"PLUGIN_SEARCH_USER_EMAIL"`
		PluginSearchUserTenantID                       string `json:"PLUGIN_SEARCH_USER_TENANT_ID"`
		PluginSearchUserGroupID                        string `json:"PLUGIN_SEARCH_USER_GROUP_ID"`
		AdminLicenseNotificationStatusSuccess          string `json:"ADMIN_LICENSE_NOTIFICATION_STATUS_SUCCESS"`
		AdminLicenseNotificationStatusFailed           string `json:"ADMIN_LICENSE_NOTIFICATION_STATUS_FAILED"`
		AdminLicenseTabKeyStored                       string `json:"ADMIN_LICENSE_TAB_KEY_STORED"`
		AdminLicenseTabKeyInvalid                      string `json:"ADMIN_LICENSE_TAB_KEY_INVALID"`
		AdminLicenseTabKeyFailed                       string `json:"ADMIN_LICENSE_TAB_KEY_FAILED"`
		AdminLicenseKey                                string `json:"ADMIN_LICENSE_KEY"`
		AdminLicenseKeyValid                           string `json:"ADMIN_LICENSE_KEY_VALID"`
		AdminLicenseCustomerID                         string `json:"ADMIN_LICENSE_CUSTOMER_ID"`
		AdminLicenseValidUntil                         string `json:"ADMIN_LICENSE_VALID_UNTIL"`
		AdminLicenseUnlimitedValidity                  string `json:"ADMIN_LICENSE_UNLIMITED_VALIDITY"`
		AdminLicenseInvalidKey                         string `json:"ADMIN_LICENSE_INVALID_KEY"`
		AdminLicenseNoKey                              string `json:"ADMIN_LICENSE_NO_KEY"`
		AdminLicenseAnotherKey                         string `json:"ADMIN_LICENSE_ANOTHER_KEY"`
		AdminLicenseYourKey                            string `json:"ADMIN_LICENSE_YOUR_KEY"`
		AdminLicensePasteKey                           string `json:"ADMIN_LICENSE_PASTE_KEY"`
		AdminLicenseSaveKey                            string `json:"ADMIN_LICENSE_SAVE_KEY"`
		AdminLicenseNotHave                            string `json:"ADMIN_LICENSE_NOT_HAVE"`
		PlgnAuditLogLabel                              string `json:"PLGN_AUDIT_LOG_LABEL"`
		PlgnAuditID                                    string `json:"PLGN_AUDIT_ID"`
		PlgnAuditEntity                                string `json:"PLGN_AUDIT_ENTITY"`
		PlgnAuditProcess                               string `json:"PLGN_AUDIT_PROCESS"`
		PlgnAuditProcessInstance                       string `json:"PLGN_AUDIT_PROCESS_INSTANCE"`
		PlgnAuditTask                                  string `json:"PLGN_AUDIT_TASK"`
		PlgnAuditCaseDefinition                        string `json:"PLGN_AUDIT_CASE_DEFINITION"`
		PlgnAuditCaseInstance                          string `json:"PLGN_AUDIT_CASE_INSTANCE"`
		PlgnAuditJobDefinition                         string `json:"PLGN_AUDIT_JOB_DEFINITION"`
		PlgnAuditJob                                   string `json:"PLGN_AUDIT_JOB"`
		PlgnAuditBatch                                 string `json:"PLGN_AUDIT_BATCH"`
		PlgnAuditDeployment                            string `json:"PLGN_AUDIT_DEPLOYMENT"`
		PlgnAuditUser                                  string `json:"PLGN_AUDIT_USER"`
		PlgnAuditTime                                  string `json:"PLGN_AUDIT_TIME"`
		PlgnAuditProperty                              string `json:"PLGN_AUDIT_PROPERTY"`
		PlgnAuditOriginalValue                         string `json:"PLGN_AUDIT_ORIGINAL_VALUE"`
		PlgnAuditNewValue                              string `json:"PLGN_AUDIT_NEW_VALUE"`
		PlgnAuditOperationsTable                       string `json:"PLGN_AUDIT_OPERATIONS_TABLE"`
		PlgnAuditOperation                             string `json:"PLGN_AUDIT_OPERATION"`
		PlgnAuditLoadMore                              string `json:"PLGN_AUDIT_LOAD_MORE"`
		PlgnAuditSearchConfigDeploymentID              string `json:"PLGN_AUDIT_SEARCH_CONFIG_DEPLOYMENT_ID"`
		PlgnAuditSearchConfigProcessDefinitionKey      string `json:"PLGN_AUDIT_SEARCH_CONFIG_PROCESS_DEFINITION_KEY"`
		PlgnAuditSearchConfigProcessInstanceID         string `json:"PLGN_AUDIT_SEARCH_CONFIG_PROCESS_INSTANCE_ID"`
		PlgnAuditSearchConfigCaseDefinitionID          string `json:"PLGN_AUDIT_SEARCH_CONFIG_CASE_DEFINITION_ID"`
		PlgnAuditSearchConfigCaseInstanceID            string `json:"PLGN_AUDIT_SEARCH_CONFIG_CASE_INSTANCE_ID"`
		PlgnAuditSearchConfigCaseExecutionID           string `json:"PLGN_AUDIT_SEARCH_CONFIG_CASE_EXECUTION_ID"`
		PlgnAuditSearchConfigTaskID                    string `json:"PLGN_AUDIT_SEARCH_CONFIG_TASK_ID"`
		PlgnAuditSearchConfigJobID                     string `json:"PLGN_AUDIT_SEARCH_CONFIG_JOB_ID"`
		PlgnAuditSearchConfigJobDefinitionID           string `json:"PLGN_AUDIT_SEARCH_CONFIG_JOB_DEFINITION_ID"`
		PlgnAuditSearchConfigUserID                    string `json:"PLGN_AUDIT_SEARCH_CONFIG_USER_ID"`
		PlgnAuditSearchOperationID                     string `json:"PLGN_AUDIT_SEARCH_OPERATION_ID"`
		PlgnAuditSearchOperationType                   string `json:"PLGN_AUDIT_SEARCH_OPERATION_TYPE"`
		PlgnAuditSearchEntityType                      string `json:"PLGN_AUDIT_SEARCH_ENTITY_TYPE"`
		PlgnAuditSearchEntityTypeIn                    string `json:"PLGN_AUDIT_SEARCH_ENTITY_TYPE_IN"`
		PlgnAuditSearchProperty                        string `json:"PLGN_AUDIT_SEARCH_PROPERTY"`
		PlgnAuditSearchAfter                           string `json:"PLGN_AUDIT_SEARCH_AFTER"`
		PlgnAuditSearchBefore                          string `json:"PLGN_AUDIT_SEARCH_BEFORE"`
		PlgnAuditSearchConfigTooltipsAddCriteria       string `json:"PLGN_AUDIT_SEARCH_CONFIG_TOOLTIPS_ADD_CRITERIA"`
		PlgnAuditSearchConfigTooltipsNotValid          string `json:"PLGN_AUDIT_SEARCH_CONFIG_TOOLTIPS_NOT_VALID"`
		PlgnAuditSearchConfigTooltipsRemoveSearch      string `json:"PLGN_AUDIT_SEARCH_CONFIG_TOOLTIPS_REMOVE_SEARCH"`
		PlgnAuditSearchConfigTooltipsType              string `json:"PLGN_AUDIT_SEARCH_CONFIG_TOOLTIPS_TYPE"`
		PlgnAuditSearchConfigTooltipsProperty          string `json:"PLGN_AUDIT_SEARCH_CONFIG_TOOLTIPS_PROPERTY"`
		PlgnAuditSearchConfigTooltipsOperator          string `json:"PLGN_AUDIT_SEARCH_CONFIG_TOOLTIPS_OPERATOR"`
		PlgnAuditSearchConfigTooltipsValue             string `json:"PLGN_AUDIT_SEARCH_CONFIG_TOOLTIPS_VALUE"`
		PlgnAuditSearchConfigProcessDefinitionID       string `json:"PLGN_AUDIT_SEARCH_CONFIG_PROCESS_DEFINITION_ID"`
		PlgnAuditSearchTimestamp                       string `json:"PLGN_AUDIT_SEARCH_TIMESTAMP"`
		PlgnAuditSearchTimestampBefore                 string `json:"PLGN_AUDIT_SEARCH_TIMESTAMP_BEFORE"`
		PlgnAuditSearchTimestampAfter                  string `json:"PLGN_AUDIT_SEARCH_TIMESTAMP_AFTER"`
		PlgnAuditNoColumns                             string `json:"PLGN_AUDIT_NO_COLUMNS"`
		PlgnAuditTextEmpty                             string `json:"PLGN_AUDIT_TEXT_EMPTY"`
		PlgnAuditEditAnnotation                        string `json:"PLGN_AUDIT_EDIT_ANNOTATION"`
		PlgnAuditAddAnnotation                         string `json:"PLGN_AUDIT_ADD_ANNOTATION"`
		PlgnAuditAnnotation                            string `json:"PLGN_AUDIT_ANNOTATION"`
		PlgnAuditEditModalHeader                       string `json:"PLGN_AUDIT_EDIT_MODAL_HEADER"`
		PlgnAuditEditNotificationSuccess               string `json:"PLGN_AUDIT_EDIT_NOTIFICATION_SUCCESS"`
		PlgnAuditEditNotificationFailure               string `json:"PLGN_AUDIT_EDIT_NOTIFICATION_FAILURE"`
		PlgnAuditDeleteModalHeader                     string `json:"PLGN_AUDIT_DELETE_MODAL_HEADER"`
		PlgnAuditDeleteNotificationSuccess             string `json:"PLGN_AUDIT_DELETE_NOTIFICATION_SUCCESS"`
		PlgnAuditDeleteNotificationFailure             string `json:"PLGN_AUDIT_DELETE_NOTIFICATION_FAILURE"`
		PlgnAuditModalDeleteConfirmation               string `json:"PLGN_AUDIT_MODAL_DELETE_CONFIRMATION"`
		PlgnSearAddColumn                              string `json:"PLGN_SEAR_ADD_COLUMN"`
		PasswordPolicyTooltip                          string `json:"PASSWORD_POLICY_TOOLTIP"`
		PasswordPolicyTooltipPartial                   string `json:"PASSWORD_POLICY_TOOLTIP_PARTIAL"`
		PasswordPolicyTooltipError                     string `json:"PASSWORD_POLICY_TOOLTIP_ERROR"`
		PasswordPolicyLength                           string `json:"PASSWORD_POLICY_LENGTH"`
		PasswordPolicyLowercase                        string `json:"PASSWORD_POLICY_LOWERCASE"`
		PasswordPolicyUppercase                        string `json:"PASSWORD_POLICY_UPPERCASE"`
		PasswordPolicyDigit                            string `json:"PASSWORD_POLICY_DIGIT"`
		PasswordPolicySpecial                          string `json:"PASSWORD_POLICY_SPECIAL"`
		PasswordPolicyUserData                         string `json:"PASSWORD_POLICY_USER_DATA"`
		FirstLoginInfo                                 string `json:"FIRST_LOGIN_INFO"`
		FirstLoginHeading                              string `json:"FIRST_LOGIN_HEADING"`
}


func (a *AdminInterface) LoadFile(jsonFile string) {
	readfile := filepath.Join("lang", jsonFile)
	dat, err := ioutil.ReadFile(readfile)
	if err != nil {
		fmt.Println("No json file: ", err)
	}
	err = json.Unmarshal(dat, a)
	if err != nil {
		fmt.Println("Unmarshall error: ", err)
	}
}

func (a *AdminInterface) SaveAdminInterface(jsonFile string, lang string) {
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
	final, err := json.MarshalIndent(a, "", "    ")
	if err != nil {
		fmt.Println("Marshall error: ", err)
	}
	defer xfile.Close()
	xfile.Write(final)
	xfile.Close()
}

func (a *AdminInterface) TranslateAdminInterface(lang string) {
	mp := make(map[string]interface{})
	frLabels := make(map[string]interface{})
	reqBodyBytes := new(bytes.Buffer)
	json.NewEncoder(reqBodyBytes).Encode(a.Labels)
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
	labs := AdminLabels{}
	json.NewEncoder(reqBodyBytes).Encode(frLabels)
	bb, err := json.Marshal(frLabels)
	err = json.Unmarshal(bb, &labs)
	a.Labels = labs
}

