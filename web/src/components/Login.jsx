import { useState } from "react"
import { useAuth } from "../utils/AuthProvider"

export default function LoginPage() {
  const { login, user } = useAuth();
  const [username, setUserName] = useState("")
  const [password, setPassword] = useState("");

  const handleSubmit = async (e) => {
    e.preventDefault();
    try {
      await login(username, password);
      // User is now logged in and stored in context
      console.log('Logged in user:', user);
    } catch (error) {
      console.error('Login failed:', error);
    }
  };
  return (
    <form onSubmit={handleSubmit}>
      <span className="inp-title"> Username </span>
      <input type="text" name="username" placeholder="johndoe" onChange={(e) => setUserName(e.target.value)} />
      <span className="inp-title"> Password </span>
      <input type="password" name="password" placeholder="*****" min={8} onChange={(e) => setPassword(e.target.value)} />
      <br />
      <button type="submit" className="button">Log In</button>
    </form>
  )
} 
