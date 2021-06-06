import { useEffect, useState, Dispatch, SetStateAction } from 'react';
import useTorrentsService from './services/useTorrentsService';
import { humanFileSize } from './utils';
import { Table, ProgressBar } from 'react-bootstrap';
import './torrents.css';
import { Torrent } from './types/Torrent';

type TorrentsProps = {
  selected?: Torrent;
  setSelected: Dispatch<SetStateAction<Torrent | undefined>>;
};

const Torrents: React.FC<TorrentsProps> = ({ selected, setSelected }) => {
  const service = useTorrentsService();

  const [_, setTime] = useState(Date.now());
  const refresh = () => {
    setTime(Date.now());
  };

  useEffect(() => {
    const interval = setInterval(() => refresh(), 1000);
    return () => {
      clearInterval(interval);
    };
  }, []);

  return (
    <div className="table-wrapper-scroll-y">
      <Table bordered hover>
        <thead>
          <tr>
            <th>Name</th>
            <th>Status</th>
            <th>Size</th>
            <th>Progress</th>
            <th>Peers</th>
          </tr>
        </thead>
        <tbody>
          {service.status === 'loading' && (
            <tr>
              <td>Loading...</td>
            </tr>
          )}
          {service.status === 'loaded' &&
            service.payload.results.map((torrent) => (
              <tr
                key={torrent.infoHash}
                onClick={() => setSelected(torrent)}
                className={`${
                  selected != null &&
                  torrent.infoHash == selected.infoHash &&
                  'table-success'
                }`}
              >
                <td>{torrent.info.name}</td>
                <td>
                  {torrent.isPaused
                    ? 'Paused'
                    : torrent.info.bytesMissing == 0
                    ? 'Seeding'
                    : 'Downloading'}
                </td>
                <td>{humanFileSize(torrent.info.length)}</td>
                <td>
                  <ProgressBar
                    now={Math.round(
                      (100 * torrent.info.bytesCompleted) / torrent.info.length
                    )}
                    label={`${Math.round(
                      (100 * torrent.info.bytesCompleted) / torrent.info.length
                    )}%`}
                    animated={!torrent.isPaused}
                    variant={
                      torrent.info.bytesMissing == 0 ? 'success' : 'info'
                    }
                  ></ProgressBar>
                </td>
                <td>{torrent.info.peers}</td>
              </tr>
            ))}
          {service.status === 'error' && (
            <div>Error reaching backend: {service.error.toString()}</div>
          )}
        </tbody>
      </Table>
    </div>
  );
};

export default Torrents;
