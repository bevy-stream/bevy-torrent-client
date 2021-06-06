import { Dispatch, SetStateAction } from "react";
import useTorrentsService from './services/useTorrentsService';
import { humanFileSize } from './utils';
import { Table, ProgressBar } from 'react-bootstrap';
import './torrents.css'
import { Torrent } from './types/Torrent'

type TorrentsProps = {
  selected?: Torrent,
  setSelected: Dispatch<SetStateAction<Torrent | undefined>>
}

const Torrents: React.FC<TorrentsProps> = ({ selected, setSelected }) => {
  const service = useTorrentsService();

  return (
    <div className="table-wrapper-scroll-y">
      <Table bordered hover>
        <thead>
          <tr>
            <th>Name</th>
            <th>Size</th>
            <th>Progress</th>
            <th>Peers</th>
          </tr>
        </thead>
        <tbody>
          {service.status === 'loading' && <tr><td>Loading...</td></tr>}
          {service.status === 'loaded' &&
            service.payload.results.map(torrent => (
              <tr
                key={torrent.infoHash}
                onClick={() => setSelected(torrent)}
                className={`${selected != null && torrent.infoHash == selected.infoHash && 'table-success'}`}
              >
                <td>{torrent.info.name}</td>
                <td>{humanFileSize(torrent.info.length)}</td>
                <td>
                  <ProgressBar
                    now={Math.round(torrent.info.bytesCompleted / torrent.info.length)}
                    label={`${Math.round(torrent.info.bytesCompleted / torrent.info.length)}%`}
                    animated={torrent.isDownloading || torrent.isUploading}
                    variant='info'
                  >
                  </ProgressBar>
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

export default Torrents
