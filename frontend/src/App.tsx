import { useEffect, useState } from 'react'
import './App.css'

type HealthResponse = {
  ok: boolean
  message: string
}

function App() {
  const [data, setData] = useState<HealthResponse | null>(null)
  const [loading, setLoading] = useState(true)
  const [error, setError] = useState<string | null>(null)

  useEffect(() => {
    let active = true

    async function loadHealth() {
      try {
        const response = await fetch('/api/health')

        if (!response.ok) {
          throw new Error(`Request failed with status ${response.status}`)
        }

        const json = (await response.json()) as HealthResponse

        if (active) {
          setData(json)
          setError(null)
        }
      } catch (err) {
        if (active) {
          setError(err instanceof Error ? err.message : 'Failed to load health')
        }
      } finally {
        if (active) {
          setLoading(false)
        }
      }
    }

    loadHealth()

    return () => {
      active = false
    }
  }, [])

  return (
    <main className="shell">
      <section className="card">
        <p className="eyebrow">Frontend + Go backend</p>
        <h1>Minimal local connection</h1>
        <p className="lead">
          The React app calls <code>/api/health</code> through the Vite proxy,
          and the Go server responds on port 8080.
        </p>

        <div className="status">
          {loading ? (
            <span className="pill">Loading backend status...</span>
          ) : error ? (
            <span className="pill pill-error">{error}</span>
          ) : (
            <span className="pill pill-success">
              {data?.ok ? 'Backend connected' : 'Backend returned unexpected data'}
            </span>
          )}
        </div>

        {data ? (
          <pre className="payload">{JSON.stringify(data, null, 2)}</pre>
        ) : null}
      </section>
    </main>
  )
}

export default App
