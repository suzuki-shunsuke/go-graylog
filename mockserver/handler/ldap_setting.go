package handler

import (
	"net/http"

	log "github.com/sirupsen/logrus"
	"github.com/suzuki-shunsuke/go-set"

	"github.com/suzuki-shunsuke/go-graylog"
	"github.com/suzuki-shunsuke/go-graylog/mockserver/logic"
	"github.com/suzuki-shunsuke/go-graylog/util"
)

// HandleGetLDAPSetting is the handler of GET LDAP Setting API.
func HandleGetLDAPSetting(
	user *graylog.User, lgc *logic.Logic, _ *http.Request, _ Params,
) (interface{}, int, error) {
	// GET /system/ldap/settings Get the LDAP configuration if it is configured
	return lgc.GetLDAPSetting()
}

// HandleUpdateLDAPSetting is the handler of Update LDAPSetting API.
func HandleUpdateLDAPSetting(
	user *graylog.User, lgc *logic.Logic, r *http.Request, ps Params,
) (interface{}, int, error) {
	// PUT /system/ldap/settings Update the LDAP configuration
	if sc, err := lgc.Authorize(user, "ldap:edit"); err != nil {
		return nil, sc, err
	}
	body, sc, err := validateRequestBody(
		r.Body, &validateReqBodyPrms{
			Required: set.NewStrSet(
				"display_name_attribute", "system_username", "search_base",
				"system_password", "ldap_uri", "search_pattern", "default_group",
			),
			Optional: set.NewStrSet(
				"enabled", "additional_default_groups", "group_mapping",
				"group_id_attribute",
				"trust_all_certificates", "use_start_tls", "group_search_base",
				"active_directory", "group_search_pattern"),
			Ignored:      nil,
			ExtForbidden: true,
		})
	if err != nil {
		return nil, sc, err
	}

	prms := &graylog.LDAPSetting{}
	if err := util.MSDecode(body, prms); err != nil {
		lgc.Logger().WithFields(log.Fields{
			"body": body, "error": err,
		}).Info("Failed to parse request body as LDAPSetting")
		return nil, 400, err
	}

	sc, err = lgc.UpdateLDAPSetting(prms)
	if err != nil {
		return nil, sc, err
	}
	if err := lgc.Save(); err != nil {
		return nil, 500, err
	}
	return nil, sc, nil
}

// HandleDeleteLDAPSetting is the handler of Delete LDAP Setting API.
func HandleDeleteLDAPSetting(
	user *graylog.User, lgc *logic.Logic, r *http.Request, ps Params,
) (interface{}, int, error) {
	// DELETE /system/ldap/settings Remove the LDAP configuration<Paste>
	// TODO authorize
	sc, err := lgc.DeleteLDAPSetting()
	if err != nil {
		return nil, sc, err
	}
	if err := lgc.Save(); err != nil {
		return nil, 500, err
	}
	return nil, sc, nil
}
