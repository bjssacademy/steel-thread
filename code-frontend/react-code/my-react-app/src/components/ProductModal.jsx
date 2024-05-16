import PropTypes from 'prop-types';
import { Modal, Button } from 'react-bootstrap';

const ProductModal = ({ product, show, onHide }) => {
  return (
    <Modal show={show} onHide={onHide}>
      <Modal.Header closeButton>
        <Modal.Title>{product.name}</Modal.Title>
      </Modal.Header>
      <Modal.Body>
        <p>Price: £{product.price}</p>
        {/* Add more details here */}
      </Modal.Body>
      <Modal.Footer>
        <Button variant="secondary" onClick={onHide}>Close</Button>
      </Modal.Footer>
    </Modal>
  );
};

ProductModal.propTypes = {
  product: PropTypes.shape({
    id: PropTypes.number.isRequired,
    name: PropTypes.string.isRequired,
    price: PropTypes.number.isRequired,
    stock_count: PropTypes.number.isRequired,
  }).isRequired,
  show: PropTypes.bool,
  onHide: PropTypes.func.isRequired,
};

export default ProductModal;
