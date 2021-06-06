import { Dispatch } from 'react';
import { humanFileSize } from './utils';
import { Table, ProgressBar } from 'react-bootstrap';
import './torrents.css';
import { Torrent } from './types/Torrent';

type TorrentsProps = {
  selectedInfoHash: string;
  setSelectedInfoHash: Dispatch<string>;
  torrents: Torrent[];
};

const Torrents: React.FC<TorrentsProps> = ({
  selectedInfoHash,
  setSelectedInfoHash,
  torrents,
}) => {
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
          {torrents.map((torrent: Torrent) => (
            <tr
              key={torrent.infoHash}
              onClick={() => setSelectedInfoHash(torrent.infoHash)}
              className={`${
                torrent.infoHash == selectedInfoHash && 'table-success'
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
                  variant={torrent.info.bytesMissing == 0 ? 'success' : 'info'}
                ></ProgressBar>
              </td>
              <td>{torrent.info.peers}</td>
            </tr>
          ))}
        </tbody>
      </Table>
    </div>
  );
};

export default Torrents;
