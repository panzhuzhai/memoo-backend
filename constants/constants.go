package constants

// ErrExitStatus represents the error status in this application.
const ErrExitStatus int = 2

const (
	// AppConfigPath is the path of application.yml.
	AppConfigPath = "resources/config/application.yml"
	// MessagesConfigPath is the path of messages.properties.
	MessagesConfigPath = "resources/config/messages.properties"
	// LoggerConfigPath is the path of zaplogger.yml.
	LoggerConfigPath       = "../resources/config/zaplogger.yml"
	JwtPrivateKeyPath      = "resources/config/jwt/private_key.pem"
	JwtPublicKeyPath       = "resources/config/jwt/public_key.pem"
	B2GameAssetContractAbi = "resources/abi/B2GameAsset.json"
	B2NFTAbi               = "resources/abi/B2NFT.json"
)

// PasswordHashCost is hash cost for a password.
const PasswordHashCost int = 10

const (
	// API represents the group of API.
	API = "/api"
	// APIBooks represents the group of book management API.
	APIBooks = API + "/books"
	// APIBooksID represents the API to get book data using id.
	APIBooksID = APIBooks + "/:id"
	// APICategories represents the group of category management API.
	APICategories = API + "/categories"
	// APIFormats represents the group of format management API.
	APIFormats = API + "/formats"
)

const (
	// APIAccount represents the group of auth management API.
	APIAccount = API + "/auth"
	// APIAccountLoginStatus represents the API to get the status of logged in account.
	APIAccountLoginStatus = APIAccount + "/loginStatus"
	// APIAccountLoginAccount represents the API to get the logged in account.
	APIAccountLoginAccount = APIAccount + "/loginAccount"
	// APIAccountLogin represents the API to login by session authentication.
	APIAccountLogin = APIAccount + "/login"
	// APIAccountLogout represents the API to logout.
	APIAccountLogout = APIAccount + "/logout"
)

const (
	// APIHealth represents the API to get the status of this application.
	APIHealth = API + "/health"
)

const (
	DataUpdateSucess    = "1"
	DataNotUpdate       = "2"
	DbErr               = "3"
	DataIllegal         = "4"
	DataIncorrectOwner  = "5"
	CollectionIncorrect = "6"
)

const (
	GAME     = "game"
	JOB      = "job"
	CONTRACT = "contract"
)

const (
	COLLECTION_LUANCHPAD = "Launchpad"
	COLLECTION_GAME      = "Game"
)

const (
	REGISTERED   = "registered"
	UNREGISTERED = "unregistered"
	CAN_CLAIM    = "can_claim"
)
