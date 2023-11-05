// Code generated by views (github.com/mpkondrashin/views). DO NOT EDIT

package smsdb

import (
    "database/sql"
    "time"
)


// ActionsetRow struct represents rows of the ACTIONSET table.
type ActionsetRow struct {
    ID	string
    Name	sql.NullString
    FlowControl	sql.NullString
    Rate	sql.NullInt32
}

// IterateActionset provide access to all rows of the ACTIONSET matching given criteria.
func IterateActionset(db *sql.DB, where string, callback func(v *ActionsetRow) error) error {
    query := "SELECT ID,NAME,FLOW_CONTROL,RATE FROM ACTIONSET"
    rows, err := db.Query(query)
    if err != nil {
        return err
    }
    defer rows.Close()
    if rows.Err() != nil {
        return rows.Err()
    }
    for rows.Next() {
        var r ActionsetRow
        err := rows.Scan(&r.ID, &r.Name, &r.FlowControl, &r.Rate)
        if err != nil {
            return err
        }
        if err := callback(&r); err != nil {
            return err
        }
    }
    return nil
}

// AlertsRow struct represents rows of the ALERTS table.
type AlertsRow struct {
    SequenceNum	uint64
    DeviceShortID	uint
    AlertTypeID	uint16
    PolicyID	string
    SignatureID	string
    BeginTime	uint64
    EndTime	uint64
    HitCount	uint
    SrcIPAddr	int64
    SrcIPAddrHigh	sql.NullInt64
    SrcPort	uint16
    DstIPAddr	int64
    DstIPAddrHigh	sql.NullInt64
    DstPort	uint16
    VirtualSegmentIndex	sql.NullInt32
    PhysicalPortIn	uint16
    VLANTag	sql.NullInt16
    Severity	uint8
    PacketTrace	uint8
    DeviceTraceBucket	uint
    DeviceTraceBeginSeq	uint
    DeviceTraceEndSeq	uint
    MessageParms	string
    Idx	uint64
    QuarantineAction	sql.NullString
    FlowControl	sql.NullString
    ActionSetUuid	sql.NullString
    ActionSetName	sql.NullString
    RateLimitRate	sql.NullString
    ClientIPAddr	[]byte
    XffIPAddr	[]byte
    TcipIPAddr	[]byte
    UriMethod	sql.NullString
    UriHost	sql.NullString
    UriString	sql.NullString
    SrcUserName	sql.NullString
    SrcDomain	sql.NullString
    SrcMachine	sql.NullString
    DstUserName	sql.NullString
    DstDomain	sql.NullString
    DstMachine	sql.NullString
}

// IterateAlerts provide access to all rows of the ALERTS matching given criteria.
func IterateAlerts(db *sql.DB, where string, callback func(v *AlertsRow) error) error {
    query := "SELECT SEQUENCE_NUM,DEVICE_SHORT_ID,ALERT_TYPE_ID,POLICY_ID,SIGNATURE_ID,BEGIN_TIME,END_TIME,HIT_COUNT,SRC_IP_ADDR,SRC_IP_ADDR_HIGH,SRC_PORT,DST_IP_ADDR,DST_IP_ADDR_HIGH,DST_PORT,VIRTUAL_SEGMENT_INDEX,PHYSICAL_PORT_IN,VLAN_TAG,SEVERITY,PACKET_TRACE,DEVICE_TRACE_BUCKET,DEVICE_TRACE_BEGIN_SEQ,DEVICE_TRACE_END_SEQ,MESSAGE_PARMS,IDX,QUARANTINE_ACTION,FLOW_CONTROL,ACTION_SET_UUID,ACTION_SET_NAME,RATE_LIMIT_RATE,CLIENT_IP_ADDR,XFF_IP_ADDR,TCIP_IP_ADDR,URI_METHOD,URI_HOST,URI_STRING,SRC_USER_NAME,SRC_DOMAIN,SRC_MACHINE,DST_USER_NAME,DST_DOMAIN,DST_MACHINE FROM ALERTS"
    rows, err := db.Query(query)
    if err != nil {
        return err
    }
    defer rows.Close()
    if rows.Err() != nil {
        return rows.Err()
    }
    for rows.Next() {
        var r AlertsRow
        err := rows.Scan(&r.SequenceNum, &r.DeviceShortID, &r.AlertTypeID, &r.PolicyID, &r.SignatureID, &r.BeginTime, &r.EndTime, &r.HitCount, &r.SrcIPAddr, &r.SrcIPAddrHigh, &r.SrcPort, &r.DstIPAddr, &r.DstIPAddrHigh, &r.DstPort, &r.VirtualSegmentIndex, &r.PhysicalPortIn, &r.VLANTag, &r.Severity, &r.PacketTrace, &r.DeviceTraceBucket, &r.DeviceTraceBeginSeq, &r.DeviceTraceEndSeq, &r.MessageParms, &r.Idx, &r.QuarantineAction, &r.FlowControl, &r.ActionSetUuid, &r.ActionSetName, &r.RateLimitRate, &r.ClientIPAddr, &r.XffIPAddr, &r.TcipIPAddr, &r.UriMethod, &r.UriHost, &r.UriString, &r.SrcUserName, &r.SrcDomain, &r.SrcMachine, &r.DstUserName, &r.DstDomain, &r.DstMachine)
        if err != nil {
            return err
        }
        if err := callback(&r); err != nil {
            return err
        }
    }
    return nil
}

// AlertTypeRow struct represents rows of the ALERT_TYPE table.
type AlertTypeRow struct {
    ID	uint64
    Name	sql.NullString
}

// IterateAlertType provide access to all rows of the ALERT_TYPE matching given criteria.
func IterateAlertType(db *sql.DB, where string, callback func(v *AlertTypeRow) error) error {
    query := "SELECT ID,NAME FROM ALERT_TYPE"
    rows, err := db.Query(query)
    if err != nil {
        return err
    }
    defer rows.Close()
    if rows.Err() != nil {
        return rows.Err()
    }
    for rows.Next() {
        var r AlertTypeRow
        err := rows.Scan(&r.ID, &r.Name)
        if err != nil {
            return err
        }
        if err := callback(&r); err != nil {
            return err
        }
    }
    return nil
}

// DeviceRow struct represents rows of the DEVICE table.
type DeviceRow struct {
    ID	string
    ShortID	uint
    Name	sql.NullString
    Model	sql.NullString
    SerialNumber	sql.NullString
    IPAddress	sql.NullString
    Location	sql.NullString
    DvVersion	sql.NullString
    OsVersion	sql.NullString
    Managed	sql.NullInt64
}

// IterateDevice provide access to all rows of the DEVICE matching given criteria.
func IterateDevice(db *sql.DB, where string, callback func(v *DeviceRow) error) error {
    query := "SELECT ID,SHORT_ID,NAME,MODEL,SERIAL_NUMBER,IP_ADDRESS,LOCATION,DV_VERSION,OS_VERSION,MANAGED FROM DEVICE"
    rows, err := db.Query(query)
    if err != nil {
        return err
    }
    defer rows.Close()
    if rows.Err() != nil {
        return rows.Err()
    }
    for rows.Next() {
        var r DeviceRow
        err := rows.Scan(&r.ID, &r.ShortID, &r.Name, &r.Model, &r.SerialNumber, &r.IPAddress, &r.Location, &r.DvVersion, &r.OsVersion, &r.Managed)
        if err != nil {
            return err
        }
        if err := callback(&r); err != nil {
            return err
        }
    }
    return nil
}

// EventCommentRow struct represents rows of the EVENT_COMMENT table.
type EventCommentRow struct {
    Idx	uint64
    Comment	sql.NullString
}

// IterateEventComment provide access to all rows of the EVENT_COMMENT matching given criteria.
func IterateEventComment(db *sql.DB, where string, callback func(v *EventCommentRow) error) error {
    query := "SELECT IDX,COMMENT FROM EVENT_COMMENT"
    rows, err := db.Query(query)
    if err != nil {
        return err
    }
    defer rows.Close()
    if rows.Err() != nil {
        return rows.Err()
    }
    for rows.Next() {
        var r EventCommentRow
        err := rows.Scan(&r.Idx, &r.Comment)
        if err != nil {
            return err
        }
        if err := callback(&r); err != nil {
            return err
        }
    }
    return nil
}

// FirewallBlockAlertsRow struct represents rows of the FIREWALL_BLOCK_ALERTS table.
type FirewallBlockAlertsRow struct {
    SequenceNum	uint64
    DeviceID	uint
    BeginTime	uint64
    EndTime	uint64
    HitCount	uint
    SrcIPAddr	uint
    SrcPort	uint16
    DstIPAddr	uint
    DstPort	uint16
    RuleID	int
    ProtocolName	sql.NullString
    ProtocolNumber	uint16
    ProtocolType	sql.NullString
    InZoneID	string
    OutZoneID	string
    PhysicalPortIn	uint16
    VLAN	sql.NullInt16
    Category	sql.NullString
    Url	sql.NullString
    UrlInfo	sql.NullString
    SeverityID	uint8
}

// IterateFirewallBlockAlerts provide access to all rows of the FIREWALL_BLOCK_ALERTS matching given criteria.
func IterateFirewallBlockAlerts(db *sql.DB, where string, callback func(v *FirewallBlockAlertsRow) error) error {
    query := "SELECT SEQUENCE_NUM,DEVICE_ID,BEGIN_TIME,END_TIME,HIT_COUNT,SRC_IP_ADDR,SRC_PORT,DST_IP_ADDR,DST_PORT,RULE_ID,PROTOCOL_NAME,PROTOCOL_NUMBER,PROTOCOL_TYPE,IN_ZONE_ID,OUT_ZONE_ID,PHYSICAL_PORT_IN,VLAN,CATEGORY,URL,URL_INFO,SEVERITY_ID FROM FIREWALL_BLOCK_ALERTS"
    rows, err := db.Query(query)
    if err != nil {
        return err
    }
    defer rows.Close()
    if rows.Err() != nil {
        return rows.Err()
    }
    for rows.Next() {
        var r FirewallBlockAlertsRow
        err := rows.Scan(&r.SequenceNum, &r.DeviceID, &r.BeginTime, &r.EndTime, &r.HitCount, &r.SrcIPAddr, &r.SrcPort, &r.DstIPAddr, &r.DstPort, &r.RuleID, &r.ProtocolName, &r.ProtocolNumber, &r.ProtocolType, &r.InZoneID, &r.OutZoneID, &r.PhysicalPortIn, &r.VLAN, &r.Category, &r.Url, &r.UrlInfo, &r.SeverityID)
        if err != nil {
            return err
        }
        if err := callback(&r); err != nil {
            return err
        }
    }
    return nil
}

// FirewallTrafficAlertsRow struct represents rows of the FIREWALL_TRAFFIC_ALERTS table.
type FirewallTrafficAlertsRow struct {
    SequenceNum	uint64
    DeviceID	uint
    EndTime	uint64
    SrcIPAddr	uint
    SrcPort	uint16
    DstIPAddr	uint
    DstPort	uint16
    RuleID	int
    ProtocolName	sql.NullString
    ProtocolNumber	uint16
    InZoneID	string
    OutZoneID	string
    Category	sql.NullString
    Duration	sql.NullString
    Url	sql.NullString
    TranferBytes	sql.NullInt64
    Message	sql.NullString
}

// IterateFirewallTrafficAlerts provide access to all rows of the FIREWALL_TRAFFIC_ALERTS matching given criteria.
func IterateFirewallTrafficAlerts(db *sql.DB, where string, callback func(v *FirewallTrafficAlertsRow) error) error {
    query := "SELECT SEQUENCE_NUM,DEVICE_ID,END_TIME,SRC_IP_ADDR,SRC_PORT,DST_IP_ADDR,DST_PORT,RULE_ID,PROTOCOL_NAME,PROTOCOL_NUMBER,IN_ZONE_ID,OUT_ZONE_ID,CATEGORY,DURATION,URL,TRANFER_BYTES,MESSAGE FROM FIREWALL_TRAFFIC_ALERTS"
    rows, err := db.Query(query)
    if err != nil {
        return err
    }
    defer rows.Close()
    if rows.Err() != nil {
        return rows.Err()
    }
    for rows.Next() {
        var r FirewallTrafficAlertsRow
        err := rows.Scan(&r.SequenceNum, &r.DeviceID, &r.EndTime, &r.SrcIPAddr, &r.SrcPort, &r.DstIPAddr, &r.DstPort, &r.RuleID, &r.ProtocolName, &r.ProtocolNumber, &r.InZoneID, &r.OutZoneID, &r.Category, &r.Duration, &r.Url, &r.TranferBytes, &r.Message)
        if err != nil {
            return err
        }
        if err := callback(&r); err != nil {
            return err
        }
    }
    return nil
}

// PolicyRow struct represents rows of the POLICY table.
type PolicyRow struct {
    ID	string
    ProfileID	string
    SignatureID	sql.NullString
    ActionsetID	sql.NullString
    Name	sql.NullString
}

// IteratePolicy provide access to all rows of the POLICY matching given criteria.
func IteratePolicy(db *sql.DB, where string, callback func(v *PolicyRow) error) error {
    query := "SELECT ID,PROFILE_ID,SIGNATURE_ID,ACTIONSET_ID,NAME FROM POLICY"
    rows, err := db.Query(query)
    if err != nil {
        return err
    }
    defer rows.Close()
    if rows.Err() != nil {
        return rows.Err()
    }
    for rows.Next() {
        var r PolicyRow
        err := rows.Scan(&r.ID, &r.ProfileID, &r.SignatureID, &r.ActionsetID, &r.Name)
        if err != nil {
            return err
        }
        if err := callback(&r); err != nil {
            return err
        }
    }
    return nil
}

// ProductCategoryRow struct represents rows of the PRODUCT_CATEGORY table.
type ProductCategoryRow struct {
    ID	sql.NullInt32
    Name	sql.NullString
    CategoryType	sql.NullString
}

// IterateProductCategory provide access to all rows of the PRODUCT_CATEGORY matching given criteria.
func IterateProductCategory(db *sql.DB, where string, callback func(v *ProductCategoryRow) error) error {
    query := "SELECT ID,NAME,CATEGORY_TYPE FROM PRODUCT_CATEGORY"
    rows, err := db.Query(query)
    if err != nil {
        return err
    }
    defer rows.Close()
    if rows.Err() != nil {
        return rows.Err()
    }
    for rows.Next() {
        var r ProductCategoryRow
        err := rows.Scan(&r.ID, &r.Name, &r.CategoryType)
        if err != nil {
            return err
        }
        if err := callback(&r); err != nil {
            return err
        }
    }
    return nil
}

// ProfileRow struct represents rows of the PROFILE table.
type ProfileRow struct {
    ID	string
    Version	string
    Name	string
    Description	sql.NullString
}

// IterateProfile provide access to all rows of the PROFILE matching given criteria.
func IterateProfile(db *sql.DB, where string, callback func(v *ProfileRow) error) error {
    query := "SELECT ID,VERSION,NAME,DESCRIPTION FROM PROFILE"
    rows, err := db.Query(query)
    if err != nil {
        return err
    }
    defer rows.Close()
    if rows.Err() != nil {
        return rows.Err()
    }
    for rows.Next() {
        var r ProfileRow
        err := rows.Scan(&r.ID, &r.Version, &r.Name, &r.Description)
        if err != nil {
            return err
        }
        if err := callback(&r); err != nil {
            return err
        }
    }
    return nil
}

// ProfileInstallInventoryRow struct represents rows of the PROFILE_INSTALL_INVENTORY table.
type ProfileInstallInventoryRow struct {
    ProfileID	string
    ProfileVersion	string
    VirtualSegmentID	sql.NullInt32
    DistributeID	string
    CompleteTime	int64
}

// IterateProfileInstallInventory provide access to all rows of the PROFILE_INSTALL_INVENTORY matching given criteria.
func IterateProfileInstallInventory(db *sql.DB, where string, callback func(v *ProfileInstallInventoryRow) error) error {
    query := "SELECT PROFILE_ID,PROFILE_VERSION,VIRTUAL_SEGMENT_ID,DISTRIBUTE_ID,COMPLETE_TIME FROM PROFILE_INSTALL_INVENTORY"
    rows, err := db.Query(query)
    if err != nil {
        return err
    }
    defer rows.Close()
    if rows.Err() != nil {
        return rows.Err()
    }
    for rows.Next() {
        var r ProfileInstallInventoryRow
        err := rows.Scan(&r.ProfileID, &r.ProfileVersion, &r.VirtualSegmentID, &r.DistributeID, &r.CompleteTime)
        if err != nil {
            return err
        }
        if err := callback(&r); err != nil {
            return err
        }
    }
    return nil
}

// QuarantineHostsRow struct represents rows of the QUARANTINE_HOSTS table.
type QuarantineHostsRow struct {
    ID	uint
    QuarantinedIP	int64
    QuarantinedIP2	int64
    QuarantinedMac	uint64
    PolicyName	sql.NullString
    State	sql.NullString
    Authority	sql.NullString
    CreateTime	time.Time
    LastUpdate	time.Time
}

// IterateQuarantineHosts provide access to all rows of the QUARANTINE_HOSTS matching given criteria.
func IterateQuarantineHosts(db *sql.DB, where string, callback func(v *QuarantineHostsRow) error) error {
    query := "SELECT ID,QUARANTINED_IP,QUARANTINED_IP_2,QUARANTINED_MAC,POLICY_NAME,STATE,AUTHORITY,CREATE_TIME,LAST_UPDATE FROM QUARANTINE_HOSTS"
    rows, err := db.Query(query)
    if err != nil {
        return err
    }
    defer rows.Close()
    if rows.Err() != nil {
        return rows.Err()
    }
    for rows.Next() {
        var r QuarantineHostsRow
        err := rows.Scan(&r.ID, &r.QuarantinedIP, &r.QuarantinedIP2, &r.QuarantinedMac, &r.PolicyName, &r.State, &r.Authority, &r.CreateTime, &r.LastUpdate)
        if err != nil {
            return err
        }
        if err := callback(&r); err != nil {
            return err
        }
    }
    return nil
}

// QuarantineNetworkDevicesRow struct represents rows of the QUARANTINE_NETWORK_DEVICES table.
type QuarantineNetworkDevicesRow struct {
    Name	sql.NullString
    IPAddress	int64
    IPAddress2	int64
}

// IterateQuarantineNetworkDevices provide access to all rows of the QUARANTINE_NETWORK_DEVICES matching given criteria.
func IterateQuarantineNetworkDevices(db *sql.DB, where string, callback func(v *QuarantineNetworkDevicesRow) error) error {
    query := "SELECT NAME,IP_ADDRESS,IP_ADDRESS_2 FROM QUARANTINE_NETWORK_DEVICES"
    rows, err := db.Query(query)
    if err != nil {
        return err
    }
    defer rows.Close()
    if rows.Err() != nil {
        return rows.Err()
    }
    for rows.Next() {
        var r QuarantineNetworkDevicesRow
        err := rows.Scan(&r.Name, &r.IPAddress, &r.IPAddress2)
        if err != nil {
            return err
        }
        if err := callback(&r); err != nil {
            return err
        }
    }
    return nil
}

// SegmentRow struct represents rows of the SEGMENT table.
type SegmentRow struct {
    ID	string
    DeviceID	string
    Name	sql.NullString
    IPAddress	sql.NullString
    SlotIndex	sql.NullInt32
    SegmentIndex	sql.NullInt32
}

// IterateSegment provide access to all rows of the SEGMENT matching given criteria.
func IterateSegment(db *sql.DB, where string, callback func(v *SegmentRow) error) error {
    query := "SELECT ID,DEVICE_ID,NAME,IP_ADDRESS,SLOT_INDEX,SEGMENT_INDEX FROM SEGMENT"
    rows, err := db.Query(query)
    if err != nil {
        return err
    }
    defer rows.Close()
    if rows.Err() != nil {
        return rows.Err()
    }
    for rows.Next() {
        var r SegmentRow
        err := rows.Scan(&r.ID, &r.DeviceID, &r.Name, &r.IPAddress, &r.SlotIndex, &r.SegmentIndex)
        if err != nil {
            return err
        }
        if err := callback(&r); err != nil {
            return err
        }
    }
    return nil
}

// SeverityRow struct represents rows of the SEVERITY table.
type SeverityRow struct {
    ID	uint
    Name	sql.NullString
}

// IterateSeverity provide access to all rows of the SEVERITY matching given criteria.
func IterateSeverity(db *sql.DB, where string, callback func(v *SeverityRow) error) error {
    query := "SELECT ID,NAME FROM SEVERITY"
    rows, err := db.Query(query)
    if err != nil {
        return err
    }
    defer rows.Close()
    if rows.Err() != nil {
        return rows.Err()
    }
    for rows.Next() {
        var r SeverityRow
        err := rows.Scan(&r.ID, &r.Name)
        if err != nil {
            return err
        }
        if err := callback(&r); err != nil {
            return err
        }
    }
    return nil
}

// SignatureRow struct represents rows of the SIGNATURE table.
type SignatureRow struct {
    ID	string
    Num	int
    Severity	int
    Name	string
    Class	string
    ProductCategoryID	int
    Protocol	string
    TaxonomyID	sql.NullInt32
    CveID	sql.NullString
    BugtraqID	sql.NullString
    Description	sql.NullString
    Message	sql.NullString
}

// IterateSignature provide access to all rows of the SIGNATURE matching given criteria.
func IterateSignature(db *sql.DB, where string, callback func(v *SignatureRow) error) error {
    query := "SELECT ID,NUM,SEVERITY,NAME,CLASS,PRODUCT_CATEGORY_ID,PROTOCOL,TAXONOMY_ID,CVE_ID,BUGTRAQ_ID,DESCRIPTION,MESSAGE FROM SIGNATURE"
    rows, err := db.Query(query)
    if err != nil {
        return err
    }
    defer rows.Close()
    if rows.Err() != nil {
        return rows.Err()
    }
    for rows.Next() {
        var r SignatureRow
        err := rows.Scan(&r.ID, &r.Num, &r.Severity, &r.Name, &r.Class, &r.ProductCategoryID, &r.Protocol, &r.TaxonomyID, &r.CveID, &r.BugtraqID, &r.Description, &r.Message)
        if err != nil {
            return err
        }
        if err := callback(&r); err != nil {
            return err
        }
    }
    return nil
}

// TaxonomyMajorRow struct represents rows of the TAXONOMY_MAJOR table.
type TaxonomyMajorRow struct {
    ID	int16
    Name	string
    Description	string
}

// IterateTaxonomyMajor provide access to all rows of the TAXONOMY_MAJOR matching given criteria.
func IterateTaxonomyMajor(db *sql.DB, where string, callback func(v *TaxonomyMajorRow) error) error {
    query := "SELECT ID,NAME,DESCRIPTION FROM TAXONOMY_MAJOR"
    rows, err := db.Query(query)
    if err != nil {
        return err
    }
    defer rows.Close()
    if rows.Err() != nil {
        return rows.Err()
    }
    for rows.Next() {
        var r TaxonomyMajorRow
        err := rows.Scan(&r.ID, &r.Name, &r.Description)
        if err != nil {
            return err
        }
        if err := callback(&r); err != nil {
            return err
        }
    }
    return nil
}

// TaxonomyMinorRow struct represents rows of the TAXONOMY_MINOR table.
type TaxonomyMinorRow struct {
    ID	int16
    MajorID	int16
    Description	string
}

// IterateTaxonomyMinor provide access to all rows of the TAXONOMY_MINOR matching given criteria.
func IterateTaxonomyMinor(db *sql.DB, where string, callback func(v *TaxonomyMinorRow) error) error {
    query := "SELECT ID,MAJOR_ID,DESCRIPTION FROM TAXONOMY_MINOR"
    rows, err := db.Query(query)
    if err != nil {
        return err
    }
    defer rows.Close()
    if rows.Err() != nil {
        return rows.Err()
    }
    for rows.Next() {
        var r TaxonomyMinorRow
        err := rows.Scan(&r.ID, &r.MajorID, &r.Description)
        if err != nil {
            return err
        }
        if err := callback(&r); err != nil {
            return err
        }
    }
    return nil
}

// TaxonomyPlatformRow struct represents rows of the TAXONOMY_PLATFORM table.
type TaxonomyPlatformRow struct {
    ID	int16
    Description	string
}

// IterateTaxonomyPlatform provide access to all rows of the TAXONOMY_PLATFORM matching given criteria.
func IterateTaxonomyPlatform(db *sql.DB, where string, callback func(v *TaxonomyPlatformRow) error) error {
    query := "SELECT ID,DESCRIPTION FROM TAXONOMY_PLATFORM"
    rows, err := db.Query(query)
    if err != nil {
        return err
    }
    defer rows.Close()
    if rows.Err() != nil {
        return rows.Err()
    }
    for rows.Next() {
        var r TaxonomyPlatformRow
        err := rows.Scan(&r.ID, &r.Description)
        if err != nil {
            return err
        }
        if err := callback(&r); err != nil {
            return err
        }
    }
    return nil
}

// TaxonomyProtocolRow struct represents rows of the TAXONOMY_PROTOCOL table.
type TaxonomyProtocolRow struct {
    ID	int16
    Description	string
}

// IterateTaxonomyProtocol provide access to all rows of the TAXONOMY_PROTOCOL matching given criteria.
func IterateTaxonomyProtocol(db *sql.DB, where string, callback func(v *TaxonomyProtocolRow) error) error {
    query := "SELECT ID,DESCRIPTION FROM TAXONOMY_PROTOCOL"
    rows, err := db.Query(query)
    if err != nil {
        return err
    }
    defer rows.Close()
    if rows.Err() != nil {
        return rows.Err()
    }
    for rows.Next() {
        var r TaxonomyProtocolRow
        err := rows.Scan(&r.ID, &r.Description)
        if err != nil {
            return err
        }
        if err := callback(&r); err != nil {
            return err
        }
    }
    return nil
}

// ThresholdUnitsRow struct represents rows of the THRESHOLD_UNITS table.
type ThresholdUnitsRow struct {
    ID	int
    Name	sql.NullString
}

// IterateThresholdUnits provide access to all rows of the THRESHOLD_UNITS matching given criteria.
func IterateThresholdUnits(db *sql.DB, where string, callback func(v *ThresholdUnitsRow) error) error {
    query := "SELECT ID,NAME FROM THRESHOLD_UNITS"
    rows, err := db.Query(query)
    if err != nil {
        return err
    }
    defer rows.Close()
    if rows.Err() != nil {
        return rows.Err()
    }
    for rows.Next() {
        var r ThresholdUnitsRow
        err := rows.Scan(&r.ID, &r.Name)
        if err != nil {
            return err
        }
        if err := callback(&r); err != nil {
            return err
        }
    }
    return nil
}

// VirtualSegmentRow struct represents rows of the VIRTUAL_SEGMENT table.
type VirtualSegmentRow struct {
    ID	uint
    DeviceID	string
    Name	sql.NullString
}

// IterateVirtualSegment provide access to all rows of the VIRTUAL_SEGMENT matching given criteria.
func IterateVirtualSegment(db *sql.DB, where string, callback func(v *VirtualSegmentRow) error) error {
    query := "SELECT ID,DEVICE_ID,NAME FROM VIRTUAL_SEGMENT"
    rows, err := db.Query(query)
    if err != nil {
        return err
    }
    defer rows.Close()
    if rows.Err() != nil {
        return rows.Err()
    }
    for rows.Next() {
        var r VirtualSegmentRow
        err := rows.Scan(&r.ID, &r.DeviceID, &r.Name)
        if err != nil {
            return err
        }
        if err := callback(&r); err != nil {
            return err
        }
    }
    return nil
}

