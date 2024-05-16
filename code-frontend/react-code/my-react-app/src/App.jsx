// App.jsx
import 'bootstrap/dist/css/bootstrap.min.css';
import { BrowserRouter, Routes, Route } from "react-router-dom";
import Layout from "./pages/Layout";
import Home from "./pages/Home";
import Blogs from "./pages/Blogs";
import Contact from "./pages/Contact";
import NoPage from "./pages/NoPage";
import Basket from "./pages/Basket";
import { useState, useEffect } from "react";
import { BasketContext } from "./context/BasketContext"; 


const App = () => {

  // eslint-disable-next-line no-unused-vars
  const [basket, setBasket] = useState([]);

  //load the basket from local storage if it exists
  useEffect(() => {
    const previousBasket = localStorage.getItem("basket");
    if (previousBasket) {
      console.log(previousBasket)
      setBasket(JSON.parse(previousBasket));
    }
  }, []);

  return (
    <BasketContext.Provider value={{ basket, setBasket }}>
      <BrowserRouter>
        <Routes>
          <Route path="/" element={<Layout />}>
            <Route index element={<Home />} />
            <Route path="blogs" element={<Blogs />} />
            <Route path="contact" element={<Contact />} />
            <Route path="basket" element={<Basket />} />
            <Route path="*" element={<NoPage />} />
          </Route>
        </Routes>
      </BrowserRouter>
    </BasketContext.Provider>
  );
};

export default App;