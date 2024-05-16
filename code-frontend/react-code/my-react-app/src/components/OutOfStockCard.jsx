import PropTypes from 'prop-types';
import { Card } from 'react-bootstrap';

const OutOfStockCard = ({ product }) => (
    <div className="col-md-4 mb-3">
    <Card style={{ height: '100%' }}>
      <Card.Img variant="top" src={`/outofstock.png`} alt={product.name} />
      <Card.Body>
        <Card.Title>{product.name}</Card.Title>
        <Card.Text>Price: £{product.price}</Card.Text>
        <Card.Text>Stock: {product.stock_count}</Card.Text>
      </Card.Body>
    </Card>
  </div>
);

OutOfStockCard.propTypes = {
    product: 
      PropTypes.shape({
        id: PropTypes.number.isRequired,
        name: PropTypes.string.isRequired,
        price: PropTypes.number.isRequired,
        stock_count: PropTypes.number.isRequired,
      }).isRequired,
  };
  
export default OutOfStockCard;