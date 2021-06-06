interface Info {
  name: string;
  bytesCompleted: number;
  bytesMissing: number;
  files: string[];
  peers: number;
  length: number;
}

export interface Torrent {
  infoHash: string;
  isUploading: boolean;
  isDownloading: boolean;
  info: Info;
}
