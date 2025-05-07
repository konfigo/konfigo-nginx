import "./App.css"

function App() {
  return (
    <>
      <h1 style={{ color: import.meta.env.VITE_COLOR }}>
        Hi {import.meta.env.VITE_NAME}
      </h1>
    </>
  )
}

export default App
