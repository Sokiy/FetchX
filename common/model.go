package common

type Config struct {
	DownloadUrl        string `toml:"download_url"`
	Total              int    `toml:"total"`
	PerPage            int    `toml:"per_page"`
	StartPage          int    `toml:"start_page"`
	DownloadDirName    string `toml:"download_dir_name"`
	DownloadSubDirName string `toml:"download_sub_dir_name"`
	FilePrefix         string `toml:"file_prefix"`
	ProcessNum         int    `toml:"process_num"`
}
