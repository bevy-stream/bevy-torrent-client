import { humanFileSize } from './utils';
import { Table, Col, Row, ProgressBar } from 'react-bootstrap';
import './torrents.css'
import { Torrent } from './types/Torrent'

type TorrentInfoProps = {
  torrent: Torrent,
}

const TorrentInfo: React.FC<TorrentInfoProps> = ({ torrent }) => {
  return (
    <Row>
      <Col className="justify-content-start text-left">
        <ProgressBar
          now={Math.round(torrent.info.bytesCompleted / torrent.info.length)}
          label={`${Math.round(torrent.info.bytesCompleted / torrent.info.length)}%`}
          animated={torrent.isDownloading || torrent.isUploading}
          variant='info'
          className="mt-3 mb-3"
        >
        </ProgressBar>

        <Row>
          <Col xs={12} sm={6} md={5} lg={4}>
            <Table className="info-table">
              <thead>
              </thead>
              <tbody>
                <tr>
                  <td><b>Down Speed:</b></td>
                  <td>TODO</td>
                </tr>
                <tr>
                  <td><b>Up Speed:</b></td>
                  <td>TODO</td>
                </tr>
                <tr>
                  <td><b>Downloaded:</b></td>
                  <td>{humanFileSize(torrent.info.bytesCompleted)}</td>
                </tr>
                <tr>
                  <td><b>Uploaded:</b></td>
                  <td>TODO</td>
                </tr>
              </tbody>
            </Table>
          </Col>
          <Col xs={12} sm={6} md={5} lg={4}>
            <Table className="info-table">
              <thead>
              </thead>
              <tbody>
                <tr>
                  <td><b>Seeds:</b></td>
                  <td>TODO</td>
                </tr>
                <tr>
                  <td><b>Peers:</b></td>
                  <td>{torrent.info.peers}</td>
                </tr>
                <tr>
                  <td><b>Share Ratio:</b></td>
                  <td>TODO</td>
                </tr>
              </tbody>
            </Table>
          </Col>
        </Row>

      </Col >
    </Row >
  );
};

export default TorrentInfo
