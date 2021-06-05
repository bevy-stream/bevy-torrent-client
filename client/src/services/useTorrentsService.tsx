import { useEffect, useState } from 'react';
import { Service } from '../types/Service';
import { Torrent } from '../types/Torrent';

export interface Torrents {
  results: Torrent[];
}

const useTorrentsService = () => {
  const [result, setResult] = useState<Service<Torrents>>({
    status: 'loading'
  });

  useEffect(() => {
    fetch('http://localhost:8080/torrents')
      .then(response => response.json())
      .then(response => setResult({ status: 'loaded', payload: response }))
      .catch(error => setResult({ status: 'error', error }));
  }, []);

  return result;
};

export default useTorrentsService;

