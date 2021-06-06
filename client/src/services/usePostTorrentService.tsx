import { useState } from 'react';
import { Service } from '../types/Service';
import { Torrent } from '../types/Torrent';

export type PostTorrent = Pick<
  Torrent,
  'isPaused' | 'infoHash'
>;

const usePostTorrentService = () => {
  const [service, setService] = useState<Service<PostTorrent>>({
    status: 'init'
  });

  const publishTorrent = (torrent: PostTorrent) => {
    setService({ status: 'loading' });

    const headers = new Headers();
    headers.append('Content-Type', 'application/json; charset=utf-8');

    return new Promise((resolve, reject) => {
      fetch(`http://localhost:8080/torrents/${torrent.infoHash}`, {
        method: 'PUT',
        body: JSON.stringify(torrent),
        headers
      })
        .then(response => response.json())
        .then(response => {
          setService({ status: 'loaded', payload: response });
          resolve(response);
        })
        .catch(error => {
          setService({ status: 'error', error });
          reject(error);
        });
    });
  };

  return {
    service,
    publishTorrent
  };
};

export default usePostTorrentService;
