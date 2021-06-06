import { useEffect, useState } from 'react';
import { Service } from '../types/Service';
import { Torrent } from '../types/Torrent';

export interface Torrents {
  results: Torrent[];
}

const useTorrentsService = () => {
  let [value, setState] = useState(0);
  // TODO: There must be a way to do this without using Date function...
  let forceUpdate = () => setState(Date.now().valueOf());

  const [result, setResult] = useState<Service<Torrents>>({
    status: 'loading',
  });

  useEffect(() => {
    fetch('http://localhost:8080/torrents')
      .then((response) => response.json())
      .then((response) => setResult({ status: 'loaded', payload: response }))
      .catch((error) => setResult({ status: 'error', error }));
  }, [value]);

  return { service: result, refreshTorrents: forceUpdate };
};

export default useTorrentsService;
