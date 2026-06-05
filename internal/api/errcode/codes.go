package errcode

// Authentication & authorization error codes
const (
	AuthUnauthorized       = "AUTH_UNAUTHORIZED"
	AuthForbidden          = "AUTH_FORBIDDEN"
	AuthLocked             = "AUTH_LOCKED"
	AuthInvalidToken       = "AUTH_INVALID_TOKEN"
	AuthExpiredToken       = "AUTH_EXPIRED_TOKEN"
	AuthInvalidCredentials = "AUTH_INVALID_CREDENTIALS"
	AuthPasswordWeak       = "AUTH_PASSWORD_WEAK"
)

// Resource error codes
const (
	ResourceNotFound    = "RESOURCE_NOT_FOUND"
	ResourceConflict    = "RESOURCE_CONFLICT"
	ResourceHasChildren = "RESOURCE_HAS_CHILDREN"
)

// IPAM error codes
const (
	SubnetOverlap      = "SUBNET_OVERLAP"
	SubnetInvalidCIDR  = "SUBNET_INVALID_CIDR"
	SubnetNotInParent  = "SUBNET_NOT_IN_PARENT"
	AddressNotInSubnet = "ADDRESS_NOT_IN_SUBNET"
	AddressOccupied    = "ADDRESS_OCCUPIED"
	NoAvailableIP      = "NO_AVAILABLE_IP"
)

// VLAN/VRF error codes
const (
	VLANIDConflict  = "VLAN_ID_CONFLICT"
	VRFNameConflict = "VRF_NAME_CONFLICT"
)

// Scan error codes
const (
	ScanTooLarge       = "SCAN_TOO_LARGE"
	ScanAlreadyRunning = "SCAN_ALREADY_RUNNING"
)

// Request error codes
const (
	RequestNotAllowed = "REQUEST_NOT_ALLOWED"
	RequestIPOccupied = "REQUEST_IP_OCCUPIED"
)

// Import/Export error codes
const (
	ImportValidationFailed = "IMPORT_VALIDATION_FAILED"
	ImportTooLarge         = "IMPORT_TOO_LARGE"
)

// General error codes
const (
	InternalError   = "INTERNAL_ERROR"
	ValidationError = "VALIDATION_ERROR"
	RackUConflict   = "RACK_U_CONFLICT"
)
