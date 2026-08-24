import { createBrowserRouter, RouterProvider } from 'react-router-dom'
import Home from './pages/Home'
import Register from './components/Register.jsx'
import Login from './components/Login.jsx'
const router = createBrowserRouter([
  {
    path: '/',
    element: <Home />,
  },
  {
    path: '/signup',
    element: <Register />,
  },
  {
    path: "/signin",
    element: <Login />
  },
])

export default function App() {
  return <RouterProvider router={router} />
}
