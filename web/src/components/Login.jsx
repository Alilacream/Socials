import { useState } from "react"

export default function LoginPage() {
  const [data, setData] = useState({ username: "", password: "" })

  const handleOnChange = (event) => {
    const { name, value } = event.target
    setData(prev => ({
      ...prev,
      [name]: value
    }));

  }

  const handleSubmit = async (event) => {
    event.preventDefault();
    const response = await fetch("http://localhost:8080/v1/login", {
      method: "POST",
      headers: {
        'Content-Type': 'application/json',
      },
      body: JSON.stringify(data),

    })
    if (!response.ok) {
      throw new Error(`HTTP error! status: ${response.status}`)
    }
    const result = response.json();
    console.log(result)
  }
  return (
    <form onSubmit={handleSubmit}>
      <span className="inp-title"> Username </span>
      <input type="text" name="username" placeholder="johndoe" onChange={handleOnChange} />
      <span className="inp-title"> Password </span>
      <input type="password" name="password" placeholder="*****" min={8} onChange={handleOnChange} />
      <button type="submit" className="button">Log In</button>
    </form>
  )
} 
