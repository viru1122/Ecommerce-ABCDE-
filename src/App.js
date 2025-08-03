import { Route, BrowserRouter as Router, Routes } from 'react-router-dom';
import ItemList from './components/ItemList';
import Login from './components/Login';

function App() {
  return (
    <Router>
      <Routes>
        <Route path="/" element={<Login />} />
        <Route path="/items" element={<ItemList />} />
      </Routes>
    </Router>
  );
}

export default App;