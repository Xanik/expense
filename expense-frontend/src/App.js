import React, { useState } from "react";
import Expenses from "./components/Expenses/Expenses";
import NewExpense from "./components/NewExpense/NewExpense";
import axios from "axios";

const App = () => {
  const [expenses, setExpenses] = useState([]);

  const fetchExpenses = React.useCallback(async () => {
    await axios
      .get(`${process.env.REACT_APP_BACKEND_HOST}/fetch`)
      .then((res) => {
        if (res.status === 200) {
          setExpenses(res.data.data.expenses);
        }
      })
      .catch((error) => {
        setExpenses([]);
      });
  }, []);

  const saveExpense = React.useCallback(async (expense) => {
    await axios
      .post(`${process.env.REACT_APP_BACKEND_HOST}/save`, {
        Title: expense.title,
        Amount: parseFloat(expense.amount),
        Date: expense.date,
      })
      .then((res) => {
        if (res.status === 200) {
          console.log("finished");

          setExpenses((prevExpenses) => {
            return [...prevExpenses, expense];
          });
        }
      })
      .catch((error) => {
        console.log({ error });
        setExpenses([]);
      });
  }, []);

  console.log({ expenses });

  React.useEffect(() => {
    fetchExpenses();
  }, [fetchExpenses]);

  const addExpenseHandler = (expense) => {
    saveExpense(expense);
  };

  return (
    <div>
      <NewExpense onAddExpense={addExpenseHandler} />
      <Expenses expenses={expenses} />
    </div>
  );
};

export default App;
