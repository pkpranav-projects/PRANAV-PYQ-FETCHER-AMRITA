<template>
  <div class="dynamic-background-container" aria-hidden="true">
    <canvas ref="canvasRef" class="topography-canvas"></canvas>
    <div class="radial-vignette"></div>
  </div>
</template>

<script setup>
import { ref, onMounted, onUnmounted } from 'vue'

const canvasRef = ref(null)
let animationFrameId = null
let isDarkMode = true

onMounted(() => {
  const canvas = canvasRef.value
  if (!canvas) return
  const ctx = canvas.getContext('2d')
  if (!ctx) return

  // Theme observer for real-time dark/light mode switching
  const updateTheme = () => {
    isDarkMode = document.documentElement.classList.contains('dark')
  }
  updateTheme()

  const observer = new MutationObserver(() => updateTheme())
  observer.observe(document.documentElement, {
    attributes: true,
    attributeFilter: ['class'],
  })

  // Viewport resize with high-DPI (Retina / 4K) scaling
  let width = 0
  let height = 0
  let dpr = 1

  const handleResize = () => {
    dpr = Math.min(window.devicePixelRatio || 1, 2)
    width = window.innerWidth
    height = window.innerHeight
    canvas.width = Math.floor(width * dpr)
    canvas.height = Math.floor(height * dpr)
    ctx.setTransform(1, 0, 0, 1, 0, 0)
    ctx.scale(dpr, dpr)
  }
  handleResize()
  window.addEventListener('resize', handleResize, { passive: true })

  // High-precision scroll velocity tracking
  let targetScrollY = window.scrollY
  let currentScrollY = window.scrollY
  let lastScrollY = window.scrollY
  let lastScrollTime = performance.now()
  let scrollVelocity = 0 // In pixels per second
  let scrollDirection = 1

  const handleScroll = () => {
    const now = performance.now()
    const currentY = window.scrollY
    const dt = Math.max(now - lastScrollTime, 8) // Prevent divide-by-zero
    const dy = currentY - lastScrollY

    if (Math.abs(dy) > 1) {
      scrollDirection = Math.sign(dy)
      const instantVelocity = (Math.abs(dy) / dt) * 1000 // px/sec
      // Fast attack: jump to peak velocity immediately
      scrollVelocity = Math.max(scrollVelocity, instantVelocity)
    }

    lastScrollY = currentY
    lastScrollTime = now
    targetScrollY = currentY
  }
  window.addEventListener('scroll', handleScroll, { passive: true })

  // Subtle interactive mouse tracking
  let mouseX = width * 0.5
  let mouseY = height * 0.5
  let targetMouseX = width * 0.5
  let targetMouseY = height * 0.5

  const handleMouseMove = (e) => {
    targetMouseX = e.clientX
    targetMouseY = e.clientY
  }
  window.addEventListener('mousemove', handleMouseMove, { passive: true })

  const prefersReducedMotion = window.matchMedia('(prefers-reduced-motion: reduce)').matches

  // Architectural topographic strata setup
  const totalRibbons = 26
  let time = 0

  // Luminescent traveling energy pulses along hero ribbons
  const energyPulses = [
    { ribbonIndex: 4,  xProgress: 0.15, baseSpeed: 0.0016, length: 220 },
    { ribbonIndex: 11, xProgress: 0.60, baseSpeed: 0.0020, length: 260 },
    { ribbonIndex: 17, xProgress: 0.35, baseSpeed: 0.0018, length: 240 },
    { ribbonIndex: 22, xProgress: 0.80, baseSpeed: 0.0022, length: 280 },
  ]

  const render = () => {
    // 1. Exponential decay of scroll velocity (smooth deceleration over ~0.6s)
    scrollVelocity *= 0.925
    if (scrollVelocity < 5) scrollVelocity = 0

    // Smooth lerp for scroll position & mouse
    currentScrollY += (targetScrollY - currentScrollY) * 0.08
    mouseX += (targetMouseX - mouseX) * 0.04
    mouseY += (targetMouseY - mouseY) * 0.04

    // Normalized velocity: 0 (idle) to 1.0 (fast scroll ~2400px/s)
    const normalizedVel = Math.min(scrollVelocity / 2400, 1.5)

    // Dynamic speed factor directly governed by scroll velocity
    // Base ambient speed is 0.003; accelerates up to 7x during fast scrolling
    const dynamicSpeedMultiplier = 1.0 + (normalizedVel * 6.5)
    const baseStep = prefersReducedMotion ? 0 : 0.003
    time += baseStep * dynamicSpeedMultiplier

    // Clear frame
    ctx.clearRect(0, 0, width, height)

    // Palette configuration for Dark and Light modes
    const colors = isDarkMode
      ? {
          primary: '1, 250, 198',       // Neon Cyan (#01FAC6)
          accent: '168, 85, 247',       // Electric Violet (#A855F7)
          sky: '56, 189, 248',          // Sky Blue (#38BDF8)
          neutral: '255, 255, 255',     // Crisp Architectural White
          crosshair: 'rgba(1, 250, 198, 0.40)',
        }
      : {
          primary: '2, 132, 199',       // Deep Cyan (#0284C7)
          accent: '99, 102, 241',       // Indigo (#6366F1)
          sky: '14, 165, 233',          // Azure (#0EA5E9)
          neutral: '15, 23, 42',        // Deep Slate (#0F172A)
          crosshair: 'rgba(2, 132, 199, 0.35)',
        }

    const scrollShift = currentScrollY * 0.0015
    const lineSpacing = height / (totalRibbons + 2)
    const normMouseY = (mouseY / (height || 1)) - 0.5

    // Slight amplitude expansion during active scroll
    const dynamicAmp = 1.0 + normalizedVel * 0.35

    // 2. Render architectural continuous topographic isolines (NO closed circles)
    for (let i = 0; i < totalRibbons; i++) {
      const ribbonProgress = i / totalRibbons
      const baseY = (i + 1.5) * lineSpacing

      // Determine ribbon hierarchy
      const isHeroCyan = i % 7 === 3
      const isHeroViolet = i % 7 === 6
      const isDashed = i % 5 === 1

      ctx.beginPath()

      if (isDashed) {
        ctx.setLineDash([5, 8])
      } else {
        ctx.setLineDash([])
      }

      // Step across canvas width
      const step = 22
      const points = []

      for (let x = -step; x <= width + step; x += step) {
        const xNorm = x / width

        // Multi-frequency harmonic topography equation
        // Wave 1: broad geographic elevation ridge
        const w1 = Math.sin(x * 0.0020 + time * 1.3 + ribbonProgress * 3.4 + scrollShift) * (28 * dynamicAmp)
        // Wave 2: medium terrain folds
        const w2 = Math.cos(x * 0.0042 - time * 0.9 + ribbonProgress * 2.2) * (15 * dynamicAmp)
        // Wave 3: fine harmonic ripple
        const w3 = Math.sin(x * 0.0085 + time * 2.2 + i * 0.45) * (5 * dynamicAmp)

        // Subtle mouse displacement
        const mDist = Math.hypot(x - mouseX, baseY - mouseY)
        const mInfluence = Math.max(0, 1 - mDist / 400) * 20 * -normMouseY

        // Elevation ridge contour modulation (naturally compresses lines on slopes)
        const ridge = Math.sin(xNorm * Math.PI * 1.6 + scrollShift * 0.7) * 20

        const y = baseY + w1 + w2 + w3 + mInfluence + ridge
        points.push({ x, y })

        if (points.length === 1) {
          ctx.moveTo(x, y)
        } else {
          const prev = points[points.length - 2]
          const midX = (prev.x + x) * 0.5
          const midY = (prev.y + y) * 0.5
          ctx.quadraticCurveTo(prev.x, prev.y, midX, midY)
        }
      }

      const lastPt = points[points.length - 1]
      ctx.lineTo(lastPt.x, lastPt.y)

      // Styling based on ribbon role
      if (isHeroCyan) {
        // Gradient hero line with vivid brand cyan
        const grad = ctx.createLinearGradient(0, 0, width, 0)
        const alpha = (isDarkMode ? 0.40 : 0.35) * (1.0 + normalizedVel * 0.25)
        grad.addColorStop(0, `rgba(${colors.primary}, 0.0)`)
        grad.addColorStop(0.2, `rgba(${colors.primary}, ${alpha * 0.9})`)
        grad.addColorStop(0.5, `rgba(${colors.sky}, ${alpha * 1.3})`)
        grad.addColorStop(0.8, `rgba(${colors.primary}, ${alpha * 0.9})`)
        grad.addColorStop(1, `rgba(${colors.primary}, 0.0)`)
        ctx.strokeStyle = grad
        ctx.lineWidth = isDarkMode ? 1.5 : 1.4
      } else if (isHeroViolet) {
        // Gradient hero line with brand violet/indigo
        const grad = ctx.createLinearGradient(0, 0, width, 0)
        const alpha = (isDarkMode ? 0.35 : 0.30) * (1.0 + normalizedVel * 0.25)
        grad.addColorStop(0, `rgba(${colors.accent}, 0.0)`)
        grad.addColorStop(0.25, `rgba(${colors.accent}, ${alpha * 0.85})`)
        grad.addColorStop(0.65, `rgba(${colors.accent}, ${alpha * 1.25})`)
        grad.addColorStop(1, `rgba(${colors.accent}, 0.0)`)
        ctx.strokeStyle = grad
        ctx.lineWidth = isDarkMode ? 1.4 : 1.3
      } else if (isDashed) {
        // Subtle technical dashed surveyor contour
        ctx.strokeStyle = `rgba(${colors.neutral}, ${isDarkMode ? 0.12 : 0.14})`
        ctx.lineWidth = 0.9
      } else {
        // Crisp neutral topographic isoline
        const baseAlpha = isDarkMode ? 0.07 : 0.09
        const depthAlpha = baseAlpha + Math.sin(ribbonProgress * Math.PI) * 0.06
        ctx.strokeStyle = `rgba(${colors.neutral}, ${depthAlpha})`
        ctx.lineWidth = 0.95
      }

      ctx.stroke()

      // Precision telemetry crosshairs (+) at harmonic intervals
      if (i % 6 === 2) {
        const checkIdx = Math.floor(points.length * (0.2 + (i % 4) * 0.18))
        if (points[checkIdx]) {
          const pt = points[checkIdx]
          const cSize = 4
          ctx.save()
          ctx.setLineDash([])
          ctx.strokeStyle = colors.crosshair
          ctx.lineWidth = 1.0
          ctx.beginPath()
          ctx.moveTo(pt.x - cSize, pt.y)
          ctx.lineTo(pt.x + cSize, pt.y)
          ctx.moveTo(pt.x, pt.y - cSize)
          ctx.lineTo(pt.x, pt.y + cSize)
          ctx.stroke()
          ctx.restore()
        }
      }
    }

    // Reset dash for pulses
    ctx.setLineDash([])

    // 3. Dynamic traveling energy pulses (accelerate proportionally with scroll speed)
    energyPulses.forEach((pulse) => {
      const pulseSpeed = pulse.baseSpeed * dynamicSpeedMultiplier
      pulse.xProgress = (pulse.xProgress + (prefersReducedMotion ? 0 : pulseSpeed)) % 1.0

      const ribbonIdx = pulse.ribbonIndex
      const baseY = (ribbonIdx + 1.5) * lineSpacing
      const centerX = pulse.xProgress * width

      const halfLen = pulse.length * 0.5
      const startX = Math.max(0, centerX - halfLen)
      const endX = Math.min(width, centerX + halfLen)

      if (startX < endX) {
        const pulseGrad = ctx.createLinearGradient(startX, 0, endX, 0)
        const pulseAlpha = (isDarkMode ? 0.65 : 0.55) * (1.0 + normalizedVel * 0.3)
        pulseGrad.addColorStop(0, `rgba(${colors.primary}, 0)`)
        pulseGrad.addColorStop(0.5, `rgba(${colors.primary}, ${pulseAlpha})`)
        pulseGrad.addColorStop(1, `rgba(${colors.primary}, 0)`)

        ctx.beginPath()
        for (let px = startX; px <= endX; px += 16) {
          const ribbonProgress = ribbonIdx / totalRibbons
          const w1 = Math.sin(px * 0.0020 + time * 1.3 + ribbonProgress * 3.4 + scrollShift) * (28 * dynamicAmp)
          const w2 = Math.cos(px * 0.0042 - time * 0.9 + ribbonProgress * 2.2) * (15 * dynamicAmp)
          const w3 = Math.sin(px * 0.0085 + time * 2.2 + ribbonIdx * 0.45) * (5 * dynamicAmp)
          const py = baseY + w1 + w2 + w3

          if (px === startX) {
            ctx.moveTo(px, py)
          } else {
            ctx.lineTo(px, py)
          }
        }

        ctx.strokeStyle = pulseGrad
        ctx.lineWidth = isDarkMode ? 2.5 : 2.2
        ctx.stroke()
      }
    })

    animationFrameId = requestAnimationFrame(render)
  }

  animationFrameId = requestAnimationFrame(render)

  onUnmounted(() => {
    if (animationFrameId) cancelAnimationFrame(animationFrameId)
    observer.disconnect()
    window.removeEventListener('resize', handleResize)
    window.removeEventListener('scroll', handleScroll)
    window.removeEventListener('mousemove', handleMouseMove)
  })
})
</script>

<style scoped>
.dynamic-background-container {
  position: fixed;
  inset: 0;
  width: 100vw;
  height: 100vh;
  pointer-events: none;
  z-index: -1;
  overflow: hidden;
  transition: background-color 0.4s ease;
}

:root:not(.dark) .dynamic-background-container {
  background: radial-gradient(circle at 50% 18%, #ffffff 0%, #f8fafc 50%, #edf2f7 100%);
}

:root.dark .dynamic-background-container {
  background: radial-gradient(circle at 50% 12%, #0b1120 0%, #060911 60%, #030508 100%);
}

.topography-canvas {
  position: absolute;
  top: 0;
  left: 0;
  width: 100%;
  height: 100%;
  display: block;
}

.radial-vignette {
  position: absolute;
  inset: 0;
  pointer-events: none;
}

:root.dark .radial-vignette {
  background: radial-gradient(circle at center, transparent 35%, rgba(3, 5, 8, 0.70) 100%);
}

:root:not(.dark) .radial-vignette {
  background: radial-gradient(circle at center, transparent 45%, rgba(226, 232, 240, 0.50) 100%);
}
</style>
