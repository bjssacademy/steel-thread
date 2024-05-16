import { useContext } from 'react';
import { Table, Button } from 'react-bootstrap';
import { BasketContext } from '../context/BasketContext';
import { PlusCircle, DashCircle, Trash } from 'react-bootstrap-icons';

const Basket = () => {
    const { basket, setBasket } = useContext(BasketContext);
    
    const calculateTotal = () => {
        let total = 0;
        for (const item of Object.values(basket)) {
          total += item.quantity * item.price;
        }
        return total;
    };
    
    const increaseQuantity = (productId) => {
        const updatedBasket = {
            ...basket,
            [productId]: {
                ...basket[productId],
                quantity: basket[productId].quantity + 1
            }
        };
        setBasket(updatedBasket);
        localStorage.setItem('basket', JSON.stringify(updatedBasket));
    };

    const decreaseQuantity = (productId) => {
        const updatedBasket = {
            ...basket,
            [productId]: {
                ...basket[productId],
                quantity: basket[productId].quantity - 1
            }
        };
        setBasket(updatedBasket);
        localStorage.setItem('basket', JSON.stringify(updatedBasket));
    };

    const removeProduct = (productId) => {
        const updatedBasket = { ...basket };
        delete updatedBasket[productId];
        setBasket(updatedBasket);
        localStorage.setItem('basket', JSON.stringify(updatedBasket));
    };

  return (
      <>
          <div className="row">
              <div className="col-md-12">
              <h3>In Your Basket</h3>
            <Table striped bordered hover>
                <thead>
                    <tr>
                    <th>Product</th>
                    <th>Quantity</th>
                    <th>Price</th>
                    <th>Subtotal</th>
                    <th></th>
                    </tr>
                </thead>
                <tbody>
                    {Object.values(basket).map(item => (
                    <tr key={item.id}>
                        <td>{item.name}</td>
                            <td>
                                <Button variant="primary" className="rounded-circle" onClick={() => increaseQuantity(item.id)}>
                                    <PlusCircle size={15} />
                                </Button> {item.quantity} <Button variant="secondary" className="rounded-circle" onClick={() => decreaseQuantity(item.id)}>
                                    <DashCircle size={15} />
                                </Button>
                            </td>
                        <td>£{item.price}</td>
                        <td>£{item.quantity * item.price}</td>
                        <td>
                            <Button variant="danger" onClick={() => removeProduct(item.id)}>
                                <Trash size={20} />
                            </Button>
                        </td>
                    </tr>
                    ))}
                </tbody>
                  </Table>
              </div>
              <div className="col-md-12"  style={{ textAlign: 'right', paddingRight: '15px' }}>
                  <hr />
                  <strong>Total: £{calculateTotal()}</strong>
              </div>
            </div>
    </>
  );
};

export default Basket;