package connection

type Connection struct {
	Redis            string            `mapstructure:"redis"`
	Rabbitmq         string            `mapstructure:"rabbitmq"`
	CoreService      ServiceConnection `mapstructure:"core_service"`
	EncodingService  ServiceConnection `mapstructure:"encoding_service"`
	UploadMp4Service ServiceConnection `mapstructure:"upload_mp4_service"`
	VideoHlsService  ServiceConnection `mapstructure:"video_hls_service"`
	QueueQuantity    string            `mapstructure:"queue-quantity"`
	Psql             PsqlConnection    `mapstructure:"psql"`
	Smpt             SmptConnection    `mapstructure:"smpt"`
}

type ServiceConnection struct {
	Host   string `mapstructure:"host"`
	Port   string `mapstructure:"port"`
	Socket string `mapstructure:"socket"`
}

type SmptConnection struct {
	Email    string `mapstructure:"email"`
	Host     string `mapstructure:"host"`
	Port     string `mapstructure:"port"`
	Password string `mapstructure:"password"`
}

type PsqlConnection struct {
	Name     string `mapstructure:"name"`
	Port     string `mapstructure:"port"`
	User     string `mapstructure:"user"`
	Password string `mapstructure:"password"`
	Host     string `mapstructure:"host"`
}
