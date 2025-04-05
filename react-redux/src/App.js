import React from 'react';
import './App.css';
import { useSelector, useDispatch } from 'react-redux';
import { increment, decrement } from './redux/counterSlice';

function App() {
  const count = useSelector((state) => state.counter.value);
  const dispatch = useDispatch();

  return (
    <div className="container">
      <h1>🔢 Redux Counter</h1>
      <div className="count">{count}</div>
      <div>
        <button onClick={() => dispatch(decrement())}>➖</button>
        <button onClick={() => dispatch(increment())}>➕</button>
      </div>
    </div>
  );
}

export default App;
