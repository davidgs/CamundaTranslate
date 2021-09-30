package main

import (
	"bytes"
	"embed"
	"encoding/json"
	"fmt"
	// "io/ioutil"
	"os"
	"strings"
	"path/filepath"


)

type Labels struct {
	APP_VENDOR                                        string `json:"APP_VENDOR,omitempty,omitempty"`
	ADMIN_APP                                         string `json:"ADMIN_APP,omitempty"`
	TASKLIST_APP                                      string `json:"TASKLIST_APP,omitempty"`
	COCKPIT_APP                                       string `json:"COCKPIT_APP,omitempty"`
	TOGGLE_NAVIGATION                                 string `json:"TOGGLE_NAVIGATION,omitempty"`
	SUBMIT                                            string `json:"SUBMIT,omitempty"`
	SAVE                                              string `json:"SAVE,omitempty"`
	START                                             string `json:"START,omitempty"`
	ABORT                                             string `json:"ABORT,omitempty"`
	CLOSE                                             string `json:"CLOSE,omitempty"`
	CANCEL                                            string `json:"CANCEL,omitempty"`
	BACK                                              string `json:"BACK,omitempty"`
	DELETE                                            string `json:"DELETE,omitempty"`
	SET                                               string `json:"SET,omitempty"`
	CHANGE                                            string `json:"CHANGE,omitempty"`
	TASK                                              string `json:"TASK,omitempty"`
	FILTER                                            string `json:"FILTER,omitempty"`
	TASKS                                             string `json:"TASKS,omitempty"`
	FILTERS                                           string `json:"FILTERS,omitempty"`
	LOADING                                           string `json:"LOADING,omitempty"`
	REQUIRED_FIELD                                    string `json:"REQUIRED_FIELD,omitempty"`
	REQUIRED_INTEGER_FIELD                            string `json:"REQUIRED_INTEGER_FIELD,omitempty"`
	REQUIRED_HEX_COLOR_FIELD                          string `json:"REQUIRED_HEX_COLOR_FIELD,omitempty"`
	PROCESS_DEFINITION                                string `json:"PROCESS_DEFINITION,omitempty"`
	PROCESS_INSTANCE                                  string `json:"PROCESS_INSTANCE,omitempty"`
	MANAGE_ACCOUNT                                    string `json:"MANAGE_ACCOUNT,omitempty"`
	ASC                                               string `json:"ASC,omitempty"`
	DESC                                              string `json:"DESC,omitempty"`
	BOOLEAN                                           string `json:"BOOLEAN,omitempty"`
	SHORT                                             string `json:"SHORT,omitempty"`
	INTEGER                                           string `json:"INTEGER,omitempty"`
	LONG                                              string `json:"LONG,omitempty"`
	DOUBLE                                            string `json:"DOUBLE,omitempty"`
	DATE                                              string `json:"DATE,omitempty"`
	STRING                                            string `json:"STRING,omitempty"`
	AUTH_LOGOUT_SUCCESSFUL                            string `json:"AUTH_LOGOUT_SUCCESSFUL,omitempty"`
	AUTH_LOGOUT_THANKS                                string `json:"AUTH_LOGOUT_THANKS,omitempty"`
	AUTH_DAY_CONTEXT_MORNING                          string `json:"AUTH_DAY_CONTEXT_MORNING,omitempty"`
	AUTH_DAY_CONTEXT_DAY                              string `json:"AUTH_DAY_CONTEXT_DAY,omitempty"`
	AUTH_DAY_CONTEXT_AFTERNOON                        string `json:"AUTH_DAY_CONTEXT_AFTERNOON,omitempty"`
	AUTH_DAY_CONTEXT_EVENING                          string `json:"AUTH_DAY_CONTEXT_EVENING,omitempty"`
	AUTH_DAY_CONTEXT_NIGHT                            string `json:"AUTH_DAY_CONTEXT_NIGHT,omitempty"`
	AUTH_DAY_CONTEXT_WEEKEND                          string `json:"AUTH_DAY_CONTEXT_WEEKEND,omitempty"`
	AUTH_FAILED_TO_DISPLAY_RESOURCE                   string `json:"AUTH_FAILED_TO_DISPLAY_RESOURCE,omitempty"`
	AUTH_AUTHENTICATION_FAILED                        string `json:"AUTH_AUTHENTICATION_FAILED,omitempty"`
	PAGE_LOGIN_ERROR_MSG                              string `json:"PAGE_LOGIN_ERROR_MSG,omitempty"`
	PAGE_LOGIN_USERNAME                               string `json:"PAGE_LOGIN_USERNAME,omitempty"`
	PAGE_LOGIN_PASSWORD                               string `json:"PAGE_LOGIN_PASSWORD,omitempty"`
	PAGE_LOGIN_PLEASE_SIGN_IN                         string `json:"PAGE_LOGIN_PLEASE_SIGN_IN,omitempty"`
	PAGE_LOGIN_SIGN_IN_ACTION                         string `json:"PAGE_LOGIN_SIGN_IN_ACTION,omitempty"`
	PAGE_LOGIN_FAILED                                 string `json:"PAGE_LOGIN_FAILED,omitempty"`
	SIGN_OUT_ACTION                                   string `json:"SIGN_OUT_ACTION,omitempty"`
	MY_PROFILE                                        string `json:"MY_PROFILE,omitempty"`
	DIRECTIVE_ENGINE_SELECT_TOOLTIP                   string `json:"DIRECTIVE_ENGINE_SELECT_TOOLTIP,omitempty"`
	DIRECTIVE_ENGINE_SELECT_STATUS_NOT_FOUND          string `json:"DIRECTIVE_ENGINE_SELECT_STATUS_NOT_FOUND,omitempty"`
	DIRECTIVE_ENGINE_SELECT_MESSAGE_NOT_FOUND         string `json:"DIRECTIVE_ENGINE_SELECT_MESSAGE_NOT_FOUND,omitempty"`
	DIRECTIVE_INPLACE_TEXTFIELD_ERROR_MSG             string `json:"DIRECTIVE_INPLACE_TEXTFIELD_ERROR_MSG,omitempty"`
	PAGES_STATUS_SERVER_ERROR                         string `json:"PAGES_STATUS_SERVER_ERROR,omitempty"`
	PAGES_MSG_SERVER_ERROR                            string `json:"PAGES_MSG_SERVER_ERROR,omitempty"`
	PAGES_STATUS_REQUEST_TIMEOUT                      string `json:"PAGES_STATUS_REQUEST_TIMEOUT,omitempty"`
	PAGES_MSG_REQUEST_TIMEOUT                         string `json:"PAGES_MSG_REQUEST_TIMEOUT,omitempty"`
	PAGES_STATUS_ACCESS_DENIED                        string `json:"PAGES_STATUS_ACCESS_DENIED,omitempty"`
	PAGES_MSG_ACCESS_DENIED                           string `json:"PAGES_MSG_ACCESS_DENIED,omitempty"`
	PAGES_MSG_ACCESS_DENIED_RESOURCE_ID               string `json:"PAGES_MSG_ACCESS_DENIED_RESOURCE_ID,omitempty"`
	PAGES_MSG_ACTION_DENIED                           string `json:"PAGES_MSG_ACTION_DENIED,omitempty"`
	PAGES_STATUS_NOT_FOUND                            string `json:"PAGES_STATUS_NOT_FOUND,omitempty"`
	PAGES_MSG_NOT_FOUND                               string `json:"PAGES_MSG_NOT_FOUND,omitempty"`
	PAGES_STATUS_COMMUNICATION_ERROR                  string `json:"PAGES_STATUS_COMMUNICATION_ERROR,omitempty"`
	PAGES_MSG_COMMUNICATION_ERROR                     string `json:"PAGES_MSG_COMMUNICATION_ERROR,omitempty"`
	PAGES_MSG_ENGINE_NOT_EXISTS                       string `json:"PAGES_MSG_ENGINE_NOT_EXISTS,omitempty"`
	SERVICES_RESOURCE_RESOLVER_ID_NOT_FOUND           string `json:"SERVICES_RESOURCE_RESOLVER_ID_NOT_FOUND,omitempty"`
	SERVICES_RESOURCE_RESOLVER_AUTH_FAILED            string `json:"SERVICES_RESOURCE_RESOLVER_AUTH_FAILED,omitempty"`
	SERVICES_RESOURCE_RESOLVER_RECEIVED_STATUS        string `json:"SERVICES_RESOURCE_RESOLVER_RECEIVED_STATUS,omitempty"`
	SERVICES_RESOURCE_RESOLVER_DISPLAY_FAILED         string `json:"SERVICES_RESOURCE_RESOLVER_DISPLAY_FAILED,omitempty"`
	SERVICES_RESOURCE_RESOLVER_NO_PROMISE             string `json:"SERVICES_RESOURCE_RESOLVER_NO_PROMISE,omitempty"`
	CAM_WIDGET_BPMN_VIEWER_ERROR                      string `json:"CAM_WIDGET_BPMN_VIEWER_ERROR,omitempty"`
	CAM_WIDGET_BPMN_VIEWER_LOADING                    string `json:"CAM_WIDGET_BPMN_VIEWER_LOADING,omitempty"`
	CAM_WIDGET_COPY                                   string `json:"CAM_WIDGET_COPY,omitempty"`
	CAM_WIDGET_COPY_LINK                              string `json:"CAM_WIDGET_COPY_LINK,omitempty"`
	CAM_WIDGET_CMMN_VIEWER_ERROR                      string `json:"CAM_WIDGET_CMMN_VIEWER_ERROR,omitempty"`
	CAM_WIDGET_CMMN_VIEWER_LOADING                    string `json:"CAM_WIDGET_CMMN_VIEWER_LOADING,omitempty"`
	CAM_WIDGET_DMN_VIEWER_ERROR                       string `json:"CAM_WIDGET_DMN_VIEWER_ERROR,omitempty"`
	CAM_WIDGET_DMN_VIEWER_LOADING                     string `json:"CAM_WIDGET_DMN_VIEWER_LOADING,omitempty"`
	CAM_WIDGET_FOOTER_TIMEZONE                        string `json:"CAM_WIDGET_FOOTER_TIMEZONE,omitempty"`
	CAM_WIDGET_FOOTER_POWERED_BY                      string `json:"CAM_WIDGET_FOOTER_POWERED_BY,omitempty"`
	CAM_WIDGET_HEADER_TOGGLE_NAVIGATION               string `json:"CAM_WIDGET_HEADER_TOGGLE_NAVIGATION,omitempty"`
	CAM_WIDGET_HEADER_MY_PROFILE                      string `json:"CAM_WIDGET_HEADER_MY_PROFILE,omitempty"`
	CAM_WIDGET_HEADER_SIGN_OUT                        string `json:"CAM_WIDGET_HEADER_SIGN_OUT,omitempty"`
	CAM_WIDGET_HEADER_SMALL_SCREEN_WARNING            string `json:"CAM_WIDGET_HEADER_SMALL_SCREEN_WARNING,omitempty"`
	CAM_WIDGET_SEARCH_MATCH_TYPE                      string `json:"CAM_WIDGET_SEARCH_MATCH_TYPE,omitempty"`
	CAM_WIDGET_SEARCH_MATCH_TYPE_ANY                  string `json:"CAM_WIDGET_SEARCH_MATCH_TYPE_ANY,omitempty"`
	CAM_WIDGET_SEARCH_MATCH_TYPE_ALL                  string `json:"CAM_WIDGET_SEARCH_MATCH_TYPE_ALL,omitempty"`
	CAM_WIDGET_SEARCH_TOTAL_NUMBER_RESULTS            string `json:"CAM_WIDGET_SEARCH_TOTAL_NUMBER_RESULTS,omitempty"`
	CAM_WIDGET_SEARCH_VARIABLE_CASE_VALUE             string `json:"CAM_WIDGET_SEARCH_VARIABLE_CASE_VALUE,omitempty"`
	CAM_WIDGET_SEARCH_VARIABLE_CASE_NAME              string `json:"CAM_WIDGET_SEARCH_VARIABLE_CASE_NAME,omitempty"`
	CAM_WIDGET_SEARCH_VARIABLE_CASE_ATTRIBUTE         string `json:"CAM_WIDGET_SEARCH_VARIABLE_CASE_ATTRIBUTE,omitempty"`
	CAM_WIDGET_VARIABLES_TABLE_TOOLTIP_SET_NULL       string `json:"CAM_WIDGET_VARIABLES_TABLE_TOOLTIP_SET_NULL,omitempty"`
	CAM_WIDGET_VARIABLES_TABLE_PLACEHOLDER_VALUE      string `json:"CAM_WIDGET_VARIABLES_TABLE_PLACEHOLDER_VALUE,omitempty"`
	CAM_WIDGET_VARIABLES_TABLE_TOOLTIP_DEFAULT_VALUE  string `json:"CAM_WIDGET_VARIABLES_TABLE_TOOLTIP_DEFAULT_VALUE,omitempty"`
	CAM_WIDGET_VARIABLES_TABLE_TOOLTIP_UPLOAD         string `json:"CAM_WIDGET_VARIABLES_TABLE_TOOLTIP_UPLOAD,omitempty"`
	CAM_WIDGET_VARIABLES_TABLE_NULL                   string `json:"CAM_WIDGET_VARIABLES_TABLE_NULL,omitempty"`
	CAM_WIDGET_VARIABLES_TABLE_EDIT_VARIABLE          string `json:"CAM_WIDGET_VARIABLES_TABLE_EDIT_VARIABLE,omitempty"`
	CAM_WIDGET_VARIABLES_TABLE_DELETE_VARIABLE        string `json:"CAM_WIDGET_VARIABLES_TABLE_DELETE_VARIABLE,omitempty"`
	CAM_WIDGET_VARIABLES_TABLE_SAVE_VARIABLE          string `json:"CAM_WIDGET_VARIABLES_TABLE_SAVE_VARIABLE,omitempty"`
	CAM_WIDGET_VARIABLES_TABLE_REVERT_VARIABLE        string `json:"CAM_WIDGET_VARIABLES_TABLE_REVERT_VARIABLE,omitempty"`
	CAM_WIDGET_VARIABLES_TABLE_NAME                   string `json:"CAM_WIDGET_VARIABLES_TABLE_NAME,omitempty"`
	CAM_WIDGET_VARIABLES_TABLE_TYPE                   string `json:"CAM_WIDGET_VARIABLES_TABLE_TYPE,omitempty"`
	CAM_WIDGET_VARIABLES_TABLE_VALUE                  string `json:"CAM_WIDGET_VARIABLES_TABLE_VALUE,omitempty"`
	CAM_WIDGET_VARIABLES_TABLE_DELETE                 string `json:"CAM_WIDGET_VARIABLES_TABLE_DELETE,omitempty"`
	CAM_WIDGET_VARIABLES_TABLE_ABORT                  string `json:"CAM_WIDGET_VARIABLES_TABLE_ABORT,omitempty"`
	CAM_WIDGET_VARIABLES_TABLE_DIALOGUE               string `json:"CAM_WIDGET_VARIABLES_TABLE_DIALOGUE,omitempty"`
	CAM_WIDGET_VARIABLE_DIALOG_INSPECT                string `json:"CAM_WIDGET_VARIABLE_DIALOG_INSPECT,omitempty"`
	CAM_WIDGET_VARIABLE_DIALOG_LABEL_OBJECT_TYPE      string `json:"CAM_WIDGET_VARIABLE_DIALOG_LABEL_OBJECT_TYPE,omitempty"`
	CAM_WIDGET_VARIABLE_DIALOG_LABEL_SERIALIZATION    string `json:"CAM_WIDGET_VARIABLE_DIALOG_LABEL_SERIALIZATION,omitempty"`
	CAM_WIDGET_VARIABLE_DIALOG_LABEL_SERIALIZED_VALUE string `json:"CAM_WIDGET_VARIABLE_DIALOG_LABEL_SERIALIZED_VALUE,omitempty"`
	CAM_WIDGET_VARIABLE_DIALOG_BTN_CLOSE              string `json:"CAM_WIDGET_VARIABLE_DIALOG_BTN_CLOSE,omitempty"`
	CAM_WIDGET_VARIABLE_DIALOG_BTN_CHANGE             string `json:"CAM_WIDGET_VARIABLE_DIALOG_BTN_CHANGE,omitempty"`
	CAM_WIDGET_VARIABLE_NULL                          string `json:"CAM_WIDGET_VARIABLE_NULL,omitempty"`
	CAM_WIDGET_VARIABLE_SET_VALUE_NULL                string `json:"CAM_WIDGET_VARIABLE_SET_VALUE_NULL,omitempty"`
	CAM_WIDGET_VARIABLE_VALUE                         string `json:"CAM_WIDGET_VARIABLE_VALUE,omitempty"`
	CAM_WIDGET_VARIABLE_UNDEFINED                     string `json:"CAM_WIDGET_VARIABLE_UNDEFINED,omitempty"`
	CAM_WIDGET_VARIABLE_RESET                         string `json:"CAM_WIDGET_VARIABLE_RESET,omitempty"`
	CAM_WIDGET_STRING_DIALOG_TITLE                    string `json:"CAM_WIDGET_STRING_DIALOG_TITLE,omitempty"`
	CAM_WIDGET_STRING_DIALOG_LABEL_VALUE              string `json:"CAM_WIDGET_STRING_DIALOG_LABEL_VALUE,omitempty"`
	CAM_WIDGET_STRING_DIALOG_LABEL_CLOSE              string `json:"CAM_WIDGET_STRING_DIALOG_LABEL_CLOSE,omitempty"`
	FILTER_CREATE                                     string `json:"FILTER_CREATE,omitempty"`
	FILTER_EDIT                                       string `json:"FILTER_EDIT,omitempty"`
	FILTER_DETAILS                                    string `json:"FILTER_DETAILS,omitempty"`
	FILTER_DELETE                                     string `json:"FILTER_DELETE,omitempty"`
	ADD_FILTER_HINT                                   string `json:"ADD_FILTER_HINT,omitempty"`
	ALL_TASKS                                         string `json:"ALL_TASKS,omitempty"`
	FILTER_DELETION_WARNING                           string `json:"FILTER_DELETION_WARNING,omitempty"`
	FILTER_DELETION_SUCCESS                           string `json:"FILTER_DELETION_SUCCESS,omitempty"`
	FILTER_DELETION_ERROR                             string `json:"FILTER_DELETION_ERROR,omitempty"`
	FILTER_SAVE_SUCCESS                               string `json:"FILTER_SAVE_SUCCESS,omitempty"`
	FILTER_SAVE_ERROR                                 string `json:"FILTER_SAVE_ERROR,omitempty"`
	FILTER_NOT_FOUND                                  string `json:"FILTER_NOT_FOUND,omitempty"`
	FILTER_FORM_BASICS                                string `json:"FILTER_FORM_BASICS,omitempty"`
	FILTER_FORM_BASICS_HINT                           string `json:"FILTER_FORM_BASICS_HINT,omitempty"`
	FILTER_NAME_LABEL                                 string `json:"FILTER_NAME_LABEL,omitempty"`
	FILTER_NAME_PLACEHOLDER                           string `json:"FILTER_NAME_PLACEHOLDER,omitempty"`
	FILTER_PRIORITY_LABEL                             string `json:"FILTER_PRIORITY_LABEL,omitempty"`
	FILTER_COLOR_LABEL                                string `json:"FILTER_COLOR_LABEL,omitempty"`
	FILTER_DESCRIPTION_LABEL                          string `json:"FILTER_DESCRIPTION_LABEL,omitempty"`
	FILTER_DESCRIPTION_PLACEHOLDER                    string `json:"FILTER_DESCRIPTION_PLACEHOLDER,omitempty"`
	FILTER_REFRESH_LABEL                              string `json:"FILTER_REFRESH_LABEL,omitempty"`
	FILTER_REFRESH_TOOLTIP                            string `json:"FILTER_REFRESH_TOOLTIP,omitempty"`
	FILTER_FORM_CRITERIA                              string `json:"FILTER_FORM_CRITERIA,omitempty"`
	FILTER_FORM_CRITERIA_HINT                         string `json:"FILTER_FORM_CRITERIA_HINT,omitempty"`
	FILTER_CRITERIA_KEY                               string `json:"FILTER_CRITERIA_KEY,omitempty"`
	FILTER_CRITERION_SUPPORT_EXPRESSION               string `json:"FILTER_CRITERION_SUPPORT_EXPRESSION,omitempty"`
	FILTER_CRITERIA_OPERATOR                          string `json:"FILTER_CRITERIA_OPERATOR,omitempty"`
	FILTER_CRITERIA_INCLUDE_ASSIGNED_TASKS            string `json:"FILTER_CRITERIA_INCLUDE_ASSIGNED_TASKS,omitempty"`
	FILTER_CRITERIA_INCLUDE_ASSIGNED_TASKS_HINT       string `json:"FILTER_CRITERIA_INCLUDE_ASSIGNED_TASKS_HINT,omitempty"`
	FILTER_CRITERIA_VALUE                             string `json:"FILTER_CRITERIA_VALUE,omitempty"`
	FILTER_ADD_CRITERION                              string `json:"FILTER_ADD_CRITERION,omitempty"`
	FILTER_REMOVE_CRITERION                           string `json:"FILTER_REMOVE_CRITERION,omitempty"`
	ADD_PERMISSION                                    string `json:"ADD_PERMISSION,omitempty"`
	FILTER_FORM_PERMISSIONS                           string `json:"FILTER_FORM_PERMISSIONS,omitempty"`
	FILTER_FORM_PERMISSIONS_GLOBAL                    string `json:"FILTER_FORM_PERMISSIONS_GLOBAL,omitempty"`
	FILTER_FORM_PERMISSIONS_LOADING_FAILURE           string `json:"FILTER_FORM_PERMISSIONS_LOADING_FAILURE,omitempty"`
	FILTER_FORM_PERMISSIONS_HINT                      string `json:"FILTER_FORM_PERMISSIONS_HINT,omitempty"`
	FILTER_FORM_PERMISSIONS_EDIT_HINT                 string `json:"FILTER_FORM_PERMISSIONS_EDIT_HINT,omitempty"`
	FILTER_FORM_PERMISSIONS_IDENTITY_TYPE_TOOLTIP     string `json:"FILTER_FORM_PERMISSIONS_IDENTITY_TYPE_TOOLTIP,omitempty"`
	FILTER_FORM_PERMISSIONS_IDENTITY_TYPE_USER        string `json:"FILTER_FORM_PERMISSIONS_IDENTITY_TYPE_USER,omitempty"`
	FILTER_FORM_PERMISSIONS_IDENTITY_TYPE_GROUP       string `json:"FILTER_FORM_PERMISSIONS_IDENTITY_TYPE_GROUP,omitempty"`
	FILTER_FORM_PERMISSIONS_REMOVE                    string `json:"FILTER_FORM_PERMISSIONS_REMOVE,omitempty"`
	FILTER_FORM_PERMISSIONS_SAVE_ERROR                string `json:"FILTER_FORM_PERMISSIONS_SAVE_ERROR,omitempty"`
	FILTER_FORM_PERMISSIONS_USER_ID                   string `json:"FILTER_FORM_PERMISSIONS_USER_ID,omitempty"`
	FILTER_FORM_PERMISSIONS_GROUP_ID                  string `json:"FILTER_FORM_PERMISSIONS_GROUP_ID,omitempty"`
	FILTER_FORM_PERMISSIONS_DUPLICATE_USER            string `json:"FILTER_FORM_PERMISSIONS_DUPLICATE_USER,omitempty"`
	FILTER_FORM_PERMISSIONS_DUPLICATE_GROUP           string `json:"FILTER_FORM_PERMISSIONS_DUPLICATE_GROUP,omitempty"`
	FILTER_FORM_PERMISSION_GROUP_USER                 string `json:"FILTER_FORM_PERMISSION_GROUP_USER,omitempty"`
	FILTER_FORM_PERMISSION_IDENTIFIER                 string `json:"FILTER_FORM_PERMISSION_IDENTIFIER,omitempty"`
	FILTER_FORM_VARIABLES                             string `json:"FILTER_FORM_VARIABLES,omitempty"`
	FILTER_FORM_VARIABLES_SHOW_UNDEFINED              string `json:"FILTER_FORM_VARIABLES_SHOW_UNDEFINED,omitempty"`
	FILTER_FORM_VARIABLES_HINT                        string `json:"FILTER_FORM_VARIABLES_HINT,omitempty"`
	FILTER_VARIABLE_NAME                              string `json:"FILTER_VARIABLE_NAME,omitempty"`
	FILTER_VARIABLE_LABEL                             string `json:"FILTER_VARIABLE_LABEL,omitempty"`
	FILTER_VARIABLE_NAME_PLACEHOLDER                  string `json:"FILTER_VARIABLE_NAME_PLACEHOLDER,omitempty"`
	FILTER_VARIABLE_LABEL_PLACEHOLDER                 string `json:"FILTER_VARIABLE_LABEL_PLACEHOLDER,omitempty"`
	FILTER_ADD_VARIABLE                               string `json:"FILTER_ADD_VARIABLE,omitempty"`
	FILTER_REMOVE_VARIABLE                            string `json:"FILTER_REMOVE_VARIABLE,omitempty"`
	FILTER_VARIABLES_AMOUNT_WARNING                   string `json:"FILTER_VARIABLES_AMOUNT_WARNING,omitempty"`
	NO_AVAILABLE_FILTER                               string `json:"NO_AVAILABLE_FILTER,omitempty"`
	FILTERS_LOADING_FAILURE                           string `json:"FILTERS_LOADING_FAILURE,omitempty"`
	INVALID_DATE                                      string `json:"INVALID_DATE,omitempty"`
	INVALID_DATE_VALUE                                string `json:"INVALID_DATE_VALUE,omitempty"`
	REQUIRE_UNIQUE_KEY                                string `json:"REQUIRE_UNIQUE_KEY,omitempty"`
	MATCH_TYPE                                        string `json:"MATCH_TYPE,omitempty"`
	MATCH_TYPE_ANY                                    string `json:"MATCH_TYPE_ANY,omitempty"`
	MATCH_TYPE_ALL                                    string `json:"MATCH_TYPE_ALL,omitempty"`
	FORM_FAILURE                                      string `json:"FORM_FAILURE,omitempty"`
	INIT_EXTERNAL_FORM_FAILED                         string `json:"INIT_EXTERNAL_FORM_FAILED,omitempty"`
	EMPTY_CONTEXT_PATH                                string `json:"EMPTY_CONTEXT_PATH,omitempty"`
	TASK_EXTERNAL_FORM_NOTE                           string `json:"TASK_EXTERNAL_FORM_NOTE,omitempty"`
	PROCESS_EXTERNAL_FORM_NOTE                        string `json:"PROCESS_EXTERNAL_FORM_NOTE,omitempty"`
	USE_GENERIC_FORM                                  string `json:"USE_GENERIC_FORM,omitempty"`
	EXTERNAL_FORM_LINK                                string `json:"EXTERNAL_FORM_LINK,omitempty"`
	SAVE_HINT                                         string `json:"SAVE_HINT,omitempty"`
	ADD_VARIABLE                                      string `json:"ADD_VARIABLE,omitempty"`
	NAME                                              string `json:"NAME,omitempty"`
	TYPE                                              string `json:"TYPE,omitempty"`
	VALUE                                             string `json:"VALUE,omitempty"`
	VARIABLE_NAME                                     string `json:"VARIABLE_NAME,omitempty"`
	VARIABLE_TYPE                                     string `json:"VARIABLE_TYPE,omitempty"`
	LOAD_VARIABLES                                    string `json:"LOAD_VARIABLES,omitempty"`
	LOAD_VARIABLES_FAILURE                            string `json:"LOAD_VARIABLES_FAILURE,omitempty"`
	BUSINESS_KEY                                      string `json:"BUSINESS_KEY,omitempty"`
	NO_TASK_VARIABLES                                 string `json:"NO_TASK_VARIABLES,omitempty"`
	REQUIRE_UNIQUE_NAME                               string `json:"REQUIRE_UNIQUE_NAME,omitempty"`
	VARIABLE_VALUE                                    string `json:"VARIABLE_VALUE,omitempty"`
	CREATE_TASK                                       string `json:"CREATE_TASK,omitempty"`
	NEW_TASK_NAME                                     string `json:"NEW_TASK_NAME,omitempty"`
	NEW_TASK_ASSIGNEE                                 string `json:"NEW_TASK_ASSIGNEE,omitempty"`
	NEW_TASK_TENANT_ID                                string `json:"NEW_TASK_TENANT_ID,omitempty"`
	NEW_TASK_DESCRIPTION                              string `json:"NEW_TASK_DESCRIPTION,omitempty"`
	TASK_SAVE_ERROR                                   string `json:"TASK_SAVE_ERROR,omitempty"`
	PROCESS_NAME                                      string `json:"PROCESS_NAME,omitempty"`
	PROCESS_START_OK                                  string `json:"PROCESS_START_OK,omitempty"`
	PROCESS_START_ERROR                               string `json:"PROCESS_START_ERROR,omitempty"`
	SEARCH_PROCESS_BY_NAME                            string `json:"SEARCH_PROCESS_BY_NAME,omitempty"`
	NO_PROCESS_DEFINITION_AVAILABLE                   string `json:"NO_PROCESS_DEFINITION_AVAILABLE,omitempty"`
	START_PROCESS                                     string `json:"START_PROCESS,omitempty"`
	CLICK_PROCESS_TO_START                            string `json:"CLICK_PROCESS_TO_START,omitempty"`
	ASSIGN_NOTE_PROCESS                               string `json:"ASSIGN_NOTE_PROCESS,omitempty"`
	ASSIGN_NOTE_CASE                                  string `json:"ASSIGN_NOTE_CASE,omitempty"`
	SHORTCUT_HELP                                     string `json:"SHORTCUT_HELP,omitempty"`
	SHORTCUT                                          string `json:"SHORTCUT,omitempty"`
	ADD                                               string `json:"ADD,omitempty"`
	REMOVE                                            string `json:"REMOVE,omitempty"`
	SUCCESSFUL                                        string `json:"SUCCESSFUL,omitempty"`
	FAILED                                            string `json:"FAILED,omitempty"`
	FAILURE                                           string `json:"FAILURE,omitempty"`
	FINISHED                                          string `json:"FINISHED,omitempty"`
	CLOSE_TASK                                        string `json:"CLOSE_TASK,omitempty"`
	TASK_HAS_BEEN_REMOVED                             string `json:"TASK_HAS_BEEN_REMOVED,omitempty"`
	INIT_GROUPS_FAILURE                               string `json:"INIT_GROUPS_FAILURE,omitempty"`
	ADD_GROUP_FAILED                                  string `json:"ADD_GROUP_FAILED,omitempty"`
	REMOVE_GROUP_FAILED                               string `json:"REMOVE_GROUP_FAILED,omitempty"`
	TASK_NAME                                         string `json:"TASK_NAME,omitempty"`
	SEARCH_BY_TASK_NAME                               string `json:"SEARCH_BY_TASK_NAME,omitempty"`
	FOLLOW_UP_DATE                                    string `json:"FOLLOW_UP_DATE,omitempty"`
	FOLLOW_UP                                         string `json:"FOLLOW_UP,omitempty"`
	DUE_DATE                                          string `json:"DUE_DATE,omitempty"`
	WITHOUT_DUE_DATE                                  string `json:"WITHOUT_DUE_DATE,omitempty"`
	DUE                                               string `json:"DUE,omitempty"`
	ASSIGNEE                                          string `json:"ASSIGNEE,omitempty"`
	ACTION                                            string `json:"ACTION,omitempty"`
	GROUP                                             string `json:"GROUP,omitempty"`
	GROUPS                                            string `json:"GROUPS,omitempty"`
	MANAGE_GROUPS                                     string `json:"MANAGE_GROUPS,omitempty"`
	GROUPS_ADD                                        string `json:"GROUPS_ADD,omitempty"`
	CURRENT_GROUPS                                    string `json:"CURRENT_GROUPS,omitempty"`
	GROUP_ID                                          string `json:"GROUP_ID,omitempty"`
	GROUP_ADD                                         string `json:"GROUP_ADD,omitempty"`
	GROUP_REMOVE                                      string `json:"GROUP_REMOVE,omitempty"`
	GROUP_ADD_ERROR                                   string `json:"GROUP_ADD_ERROR,omitempty"`
	GROUP_DELETE_ERROR                                string `json:"GROUP_DELETE_ERROR,omitempty"`
	ADD_GROUPS                                        string `json:"ADD_GROUPS,omitempty"`
	USE_ADD_GROUP                                     string `json:"USE_ADD_GROUP,omitempty"`
	DUPLICATE_GROUP                                   string `json:"DUPLICATE_GROUP,omitempty"`
	OWNER                                             string `json:"OWNER,omitempty"`
	DIAGRAM                                           string `json:"DIAGRAM,omitempty"`
	RENDERING                                         string `json:"RENDERING,omitempty"`
	DIAGRAM_RENDER_FAILURE                            string `json:"DIAGRAM_RENDER_FAILURE,omitempty"`
	FORM                                              string `json:"FORM,omitempty"`
	NO_TASK_FORM                                      string `json:"NO_TASK_FORM,omitempty"`
	HISTORY                                           string `json:"HISTORY,omitempty"`
	NO_HISTORY                                        string `json:"NO_HISTORY,omitempty"`
	COMMENT                                           string `json:"COMMENT,omitempty"`
	COMMENT_CREATE                                    string `json:"COMMENT_CREATE,omitempty"`
	COMMENT_MESSAGE_PLACEHOLDER                       string `json:"COMMENT_MESSAGE_PLACEHOLDER,omitempty"`
	COMMENT_SAVE_ERROR                                string `json:"COMMENT_SAVE_ERROR,omitempty"`
	DESCRIPTION                                       string `json:"DESCRIPTION,omitempty"`
	NO_DESCRIPTION                                    string `json:"NO_DESCRIPTION,omitempty"`
	NO_DIAGRAM                                        string `json:"NO_DIAGRAM,omitempty"`
	SELECT_TASK_IN_LIST                               string `json:"SELECT_TASK_IN_LIST,omitempty"`
	CLAIM_TASK                                        string `json:"CLAIM_TASK,omitempty"`
	UNCLAIM_TASK                                      string `json:"UNCLAIM_TASK,omitempty"`
	RESET_TASK_ASSIGNEE                               string `json:"RESET_TASK_ASSIGNEE,omitempty"`
	DELEGATE_TASK                                     string `json:"DELEGATE_TASK,omitempty"`
	SET_FOLLOW_UP_DATE                                string `json:"SET_FOLLOW_UP_DATE,omitempty"`
	RESET_FOLLOW_UP_DATE                              string `json:"RESET_FOLLOW_UP_DATE,omitempty"`
	SET_DUE_DATE                                      string `json:"SET_DUE_DATE,omitempty"`
	RESET_DUE_DATE                                    string `json:"RESET_DUE_DATE,omitempty"`
	ASSIGNED_TO_YOU                                   string `json:"ASSIGNED_TO_YOU,omitempty"`
	OWNER_IS_YOU                                      string `json:"OWNER_IS_YOU,omitempty"`
	ASSIGNEE_RESET_OK                                 string `json:"ASSIGNEE_RESET_OK,omitempty"`
	ASSIGNED_RESET_ERROR                              string `json:"ASSIGNEE_RESET_ERROR,omitempty"`
	INSTANCE_LINK                                     string `json:"INSTANCE_LINK,omitempty"`
	TASK_NO_PROCESS_DEFINITION                        string `json:"TASK_NO_PROCESS_DEFINITION,omitempty"`
	CLAIM_OK                                          string `json:"CLAIM_OK,omitempty"`
	CLAIM_ERROR                                       string `json:"CLAIM_ERROR,omitempty"`
	UNCLAIM_OK                                        string `json:"UNCLAIM_OK,omitempty"`
	UNCLAIM_ERROR                                     string `json:"UNCLAIM_ERROR,omitempty"`
	RESOLVE_OK                                        string `json:"RESOLVE_OK,omitempty"`
	RESOLVE_ERROR                                     string `json:"RESOLVE_ERROR,omitempty"`
	DELEGATE_ERROR                                    string `json:"DELEGATE_ERROR,omitempty"`
	DELEGATE_OK                                       string `json:"DELEGATE_OK,omitempty"`
	ASSIGNED_ERROR                                    string `json:"ASSIGNED_ERROR,omitempty"`
	ASSIGNED_OK                                       string `json:"ASSIGNED_OK,omitempty"`
	COMPLETE_ERROR                                    string `json:"COMPLETE_ERROR,omitempty"`
	COMPLETE_OK                                       string `json:"COMPLETE_OK,omitempty"`
	TASK_UPDATE_SUCESS                                string `json:"TASK_UPDATE_SUCESS,omitempty"`
	TASK_UPDATE_ERROR                                 string `json:"TASK_UPDATE_ERROR,omitempty"`
	TASK_NOT_EXIST                                    string `json:"TASK_NOT_EXIST,omitempty"`
	INSTANCE_SUSPENDED                                string `json:"INSTANCE_SUSPENDED,omitempty"`
	COMPLETE                                          string `json:"COMPLETE,omitempty"`
	NONE                                              string `json:"NONE,omitempty"`
	Assign                                            string `json:"Assign,omitempty"`
	Claim                                             string `json:"Claim,omitempty"`
	Complete                                          string `json:"Complete,omitempty"`
	Create                                            string `json:"Create,omitempty"`
	Delegate                                          string `json:"Delegate,omitempty"`
	Delete                                            string `json:"Delete,omitempty"`
	Resolve                                           string `json:"Resolve,omitempty"`
	SetOwner                                          string `json:"SetOwner,omitempty"`
	SetPriority                                       string `json:"SetPriority,omitempty"`
	Update                                            string `json:"Update,omitempty"`
	Activate                                          string `json:"Activate,omitempty"`
	Suspend                                           string `json:"Suspend,omitempty"`
	AddUserLink                                       string `json:"AddUserLink,omitempty"`
	DeleteUserLink                                    string `json:"DeleteUserLink,omitempty"`
	AddGroupLink                                      string `json:"AddGroupLink,omitempty"`
	DeleteGroupLink                                   string `json:"DeleteGroupLink,omitempty"`
	AddAttachment                                     string `json:"AddAttachment,omitempty"`
	DeleteAttachment                                  string `json:"DeleteAttachment,omitempty"`
	Assignee                                          string `json:"assignee,omitempty"`
	Delegation                                        string `json:"delegation,omitempty"`
	Delete_                                           string `json:"delete,omitempty"`
	Description                                       string `json:"description,omitempty"`
	DueDate                                           string `json:"dueDate,omitempty"`
	FollowUpDate                                      string `json:"followUpDate,omitempty"`
	Name                                              string `json:"name,omitempty"`
	Owner                                             string `json:"owner,omitempty"`
	ParentTask                                        string `json:"parentTask,omitempty"`
	Priority                                          string `json:"priority,omitempty"`
	CaseInstanceId                                    string `json:"caseInstanceId,omitempty"`
	Comment                                           string `json:"Comment,omitempty"`
	SORT_BY                                           string `json:"SORT_BY,omitempty"`
	ADD_SORT_BY                                       string `json:"ADD_SORT_BY,omitempty"`
	REMOVE_SORTING                                    string `json:"REMOVE_SORTING,omitempty"`
	PRIORITY                                          string `json:"PRIORITY,omitempty"`
	CREATION_DATE                                     string `json:"CREATION_DATE,omitempty"`
	CREATION                                          string `json:"CREATION,omitempty"`
	SEARCH_PLACEHOLDER                                string `json:"SEARCH_PLACEHOLDER,omitempty"`
	DELETE_SEARCH                                     string `json:"DELETE_SEARCH,omitempty"`
	PROPERTY                                          string `json:"PROPERTY,omitempty"`
	OPERATOR                                          string `json:"OPERATOR,omitempty"`
	INVALID_SEARCH                                    string `json:"INVALID_SEARCH,omitempty"`
	BEFORE                                            string `json:"BEFORE,omitempty"`
	AFTER                                             string `json:"AFTER,omitempty"`
	SEARCH_TASK_BY_NAME                               string `json:"SEARCH_TASK_BY_NAME,omitempty"`
	NO_MATCHING_TASK                                  string `json:"NO_MATCHING_TASK,omitempty"`
	RESTRICTION_NOTICE                                string `json:"RESTRICTION_NOTICE,omitempty"`
	RESET_PAGE                                        string `json:"RESET_PAGE,omitempty"`
	TASK_LIST_LOADING_FAILURE                         string `json:"TASK_LIST_LOADING_FAILURE,omitempty"`
	PROCESS_VARIABLE                                  string `json:"PROCESS_VARIABLE,omitempty"`
	TASK_VARIABLE                                     string `json:"TASK_VARIABLE,omitempty"`
	CASE_VARIABLE                                     string `json:"CASE_VARIABLE,omitempty"`
	EXECUTION_VARIABLE                                string `json:"EXECUTION_VARIABLE,omitempty"`
	CASE_EXECUTION_VARIABLE                           string `json:"CASE_EXECUTION_VARIABLE,omitempty"`
	CASE_INSTANCE_VARIABLE                            string `json:"CASE_INSTANCE_VARIABLE,omitempty"`
	PROCESS_INSTANCE_ID                               string `json:"PROCESS_INSTANCE_ID,omitempty"`
	PROCESS_INSTANCE_BUSINESS_KEY                     string `json:"PROCESS_INSTANCE_BUSINESS_KEY,omitempty"`
	PROCESS_DEFINITION_ID                             string `json:"PROCESS_DEFINITION_ID,omitempty"`
	PROCESS_DEFINITION_DEFINITION_KEY                 string `json:"PROCESS_DEFINITION_KEY,omitempty"`
	PROCESS_DEFINITION_NAME                           string `json:"PROCESS_DEFINITION_NAME,omitempty"`
	EXECUTION_ID                                      string `json:"EXECUTION_ID,omitempty"`
	CASE_INSTANCE_ID                                  string `json:"CASE_INSTANCE_ID,omitempty"`
	CASE_INSTANCE_BUSINESS_KEY                        string `json:"CASE_INSTANCE_BUSINESS_KEY,omitempty"`
	CASE_DEFINITION_ID                                string `json:"CASE_DEFINITION_ID,omitempty"`
	CASE_DEFINITION_KEY                               string `json:"CASE_DEFINITION_KEY,omitempty"`
	CASE_DEFINITION_NAME                              string `json:"CASE_DEFINITION_NAME,omitempty"`
	CASE_EXECUTION_ID                                 string `json:"CASE_EXECUTION_ID,omitempty"`
	CANDIDATE_GROUP                                   string `json:"CANDIDATE_GROUP,omitempty"`
	CANDIDATE_USER                                    string `json:"CANDIDATE_USER,omitempty"`
	INVOLVED_USER                                     string `json:"INVOLVED_USER,omitempty"`
	TASK_DEFINITION_KEY                               string `json:"TASK_DEFINITION_KEY,omitempty"`
	DELEGATION_STATE                                  string `json:"DELEGATION_STATE,omitempty"`
	LIKE                                              string `json:"LIKE,omitempty"`
	TENANT_ID                                         string `json:"TENANT_ID,omitempty"`
	WITHOUT_TENANT_ID                                 string `json:"WITHOUT_TENANT_ID,omitempty"`
	EMPTY_VALUE                                       string `json:"EMPTY_VALUE,omitempty"`
	TASK_VARIABLES                                    string `json:"TASK_VARIABLES,omitempty"`
	VARIABLE_OBJECT_TYPE_NAME                         string `json:"VARIABLE_OBJECT_TYPE_NAME,omitempty"`
	VARIABLE_VALUE_INFO                               string `json:"VARIABLE_VALUE_INFO,omitempty"`
	VARIABLE_SERIALIZATION_DATA_FORMAT                string `json:"VARIABLE_SERIALIZATION_DATA_FORMAT,omitempty"`
	DOWNLOAD                                          string `json:"DOWNLOAD,omitempty"`
	SHOW_LESS                                         string `json:"SHOW_LESS,omitempty"`
	SHOW_MORE                                         string `json:"SHOW_MORE,omitempty"`
	UNDEFINED_VARIABLE                                string `json:"UNDEFINED_VARIABLE,omitempty"`
	TOGGLE                                            string `json:"TOGGLE,omitempty"`
	SERIALIZED                                        string `json:"SERIALIZED,omitempty"`
	DESERIALIZED                                      string `json:"DESERIALIZED,omitempty"`
	DESERIALIZATION_ERROR                             string `json:"DESERIALIZATION_ERROR,omitempty"`
	FIRST_LOGIN_INFO                                  string `json:"FIRST_LOGIN_INFO,omitempty"`
	FIRST_LOGIN_HEADING                               string `json:"FIRST_LOGIN_HEADING,omitempty"`
	TELEMETRY_SAVE                                    string `json:"TELEMETRY_SAVE,omitempty"`
	TELEMETRY_CANCEL                                  string `json:"TELEMETRY_CANCEL,omitempty"`
	TELEMETRY_SETTINGS                                string `json:"TELEMETRY_SETTINGS,omitempty"`
	TELEMETRY_INTRODUCTION                            string `json:"TELEMETRY_INTRODUCTION,omitempty"`
	TELEMETRY_ENABLE                                  string `json:"TELEMETRY_ENABLE,omitempty"`
	TELEMETRY_DETAILS                                 string `json:"TELEMETRY_DETAILS,omitempty"`
	TELEMETRY_PRIVACY                                 string `json:"TELEMETRY_PRIVACY,omitempty"`
	TELEMETRY_DATA                                    string `json:"TELEMETRY_DATA,omitempty"`
	TELEMETRY_DOCUMENTATION                           string `json:"TELEMETRY_DOCUMENTATION,omitempty"`
	TELEMETRY_OR_VIEW                                 string `json:"TELEMETRY_OR_VIEW,omitempty"`
	TELEMETRY_PRIVACY_POLICY                          string `json:"TELEMETRY_PRIVACY_POLICY,omitempty"`
	TELEMETRY_SUCCESS_HEADING                         string `json:"TELEMETRY_SUCCESS_HEADING,omitempty"`
	TELEMETRY_SUCCESS                                 string `json:"TELEMETRY_SUCCESS,omitempty"`
	TELEMETRY_ERROR_STATUS                            string `json:"TELEMETRY_ERROR_STATUS,omitempty"`
	TELEMETRY_ERROR_MESSAGE                           string `json:"TELEMETRY_ERROR_MESSAGE,omitempty"`
}
type TaskListInterface struct {
	Labels Labels      `json:"labels,omitempty"`
	Dates  DateLocales `json:"dateLocales,omitempty"`
}

func (t *TaskListInterface) LoadFile(jsonFile string, dir embed.FS) {
	readfile := filepath.Join("lang", jsonFile)
	dat, err := dir.ReadFile(readfile)
//	dat, err := ioutil.ReadFile(readfile)
	if err != nil {
		fmt.Println("No json file: ", err)
	}
	err = json.Unmarshal(dat, t)
	if err != nil {
		fmt.Println("Unmarshall error: ", err)
	}
}

func (t *TaskListInterface) SaveTaskListInterface(jsonFile string, lang string) {
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
	final, err := json.MarshalIndent(t, "", "    ")
	if err != nil {
		fmt.Println("Marshall error: ", err)
	}
	defer xfile.Close()
	xfile.Write(final)
	xfile.Close()
}

func (t *TaskListInterface) TranslateTaskListInterface(lang string) {
	mp := make(map[string]interface{})
	frLabels := make(map[string]interface{})
	reqBodyBytes := new(bytes.Buffer)
	json.NewEncoder(reqBodyBytes).Encode(t.Labels)
	err := json.Unmarshal(reqBodyBytes.Bytes(), &mp)
	if err != nil {
		fmt.Println("Unmarshall error: ", err)
	}
	translator := Translator{}

	for k, v := range mp {
		s, err := translator.TranslateTextWithModel(lang, fmt.Sprintf("%v", v), "nmt")
		if err != nil {
			fmt.Println("Translate Error: ", err)
		}
		frLabels[k] = s
	}
	labs := Labels{}
	json.NewEncoder(reqBodyBytes).Encode(frLabels)
	bb, err := json.Marshal(frLabels)
	err = json.Unmarshal(bb, &labs)
	t.Labels = labs
}


