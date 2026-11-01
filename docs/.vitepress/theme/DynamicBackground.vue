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

  // Real-time dark/light mode observer
  const updateTheme = () => {
    isDarkMode = document.documentElement.classList.contains('dark')
  }
  updateTheme()

  const observer = new MutationObserver(() => updateTheme())
  observer.observe(document.documentElement, {
    attributes: true,
    attributeFilter: ['class'],
  })

  // High-DPI Canvas scaling
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

  // Fluid scroll tracking with hydrodynamic momentum
  let targetScrollY = window.scrollY
  let currentScrollY = window.scrollY
  let lastScrollY = window.scrollY
  let lastScrollTime = performance.now()
  let scrollVelocity = 0
  let verticalFlowMomentum = 0
  let continuousVerticalOffset = 0

  const handleScroll = () => {
    const now = performance.now()
    const currentY = window.scrollY
    const dt = Math.max(now - lastScrollTime, 8)
    const dy = currentY - lastScrollY

    if (Math.abs(dy) > 0.5) {
      const instantVelocity = (Math.abs(dy) / dt) * 1000
      // Fast attack
      scrollVelocity = Math.max(scrollVelocity, instantVelocity)
      // Inject physical momentum into vertical flow stream (smooth flow speed)
      verticalFlowMomentum += dy * 0.42
    }

    lastScrollY = currentY
    lastScrollTime = now
    targetScrollY = currentY
  }
  window.addEventListener('scroll', handleScroll, { passive: true })

  // Interactive mouse influence
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

  // Architectural contour configuration
  const totalRibbons = 26
  let time = 0

  // Luminescent tracer pulses
  const energyPulses = [
    { ribbonIndex: 3,  xProgress: 0.10, speed: 0.0022, length: 240 },
    { ribbonIndex: 9,  xProgress: 0.50, speed: 0.0026, length: 280 },
    { ribbonIndex: 16, xProgress: 0.30, speed: 0.0020, length: 260 },
    { ribbonIndex: 21, xProgress: 0.75, speed: 0.0024, length: 300 },
  ]

  const render = () => {
    // 1. Hydrodynamic decay of scroll velocity and vertical flow momentum
    scrollVelocity *= 0.93
    if (scrollVelocity < 5) scrollVelocity = 0

    // Smooth fluid cushioning on vertical flow momentum
    verticalFlowMomentum *= 0.88

    // Ambient fluid drift speed + scroll momentum
    const ambientFlow = prefersReducedMotion ? 0 : 0.45
    const dynamicFlowStep = ambientFlow + (verticalFlowMomentum * 0.14)
    continuousVerticalOffset += dynamicFlowStep

    // Lerp positions
    currentScrollY += (targetScrollY - currentScrollY) * 0.08
    mouseX += (targetMouseX - mouseX) * 0.05
    mouseY += (targetMouseY - mouseY) * 0.05

    // Normalized velocity for dynamic responsiveness
    const normVel = Math.min(scrollVelocity / 1800, 1.8)
    const speedBoost = 1.0 + (normVel * 5.0)
    const baseStep = prefersReducedMotion ? 0 : 0.004
    time += baseStep * speedBoost

    // Dynamic wave amplitude expansion during active scrolling
    const dynamicAmp = 1.0 + (normVel * 0.45)

    // Clear frame
    ctx.clearRect(0, 0, width, height)

    // Professional refined color palette (subtle engineering sapphire, slate, crisp monochrome)
    const colors = isDarkMode
      ? {
          primary: '56, 189, 248',       // Refined Electric Sky (#38BDF8)
          accent: '129, 140, 248',      // Sophisticated Indigo (#818CF8)
          highlight: '241, 245, 249',   // Crisp Titanium White
          neutral: '148, 163, 184',     // Architectural Slate (#94A3B8)
          crosshair: 'rgba(56, 189, 248, 0.55)',
        }
      : {
          primary: '2, 132, 199',       // Deep Technical Azure (#0284C7)
          accent: '79, 70, 229',        // Corporate Indigo (#4F46E5)
          highlight: '15, 23, 42',      // Deep Obsidian Slate
          neutral: '71, 85, 105',       // Technical Slate (#475569)
          crosshair: 'rgba(2, 132, 199, 0.45)',
        }

    // Canvas vertical looping boundaries for continuous seamless stream
    const padding = 160
    const totalSpan = height + padding * 2
    const spacing = totalSpan / totalRibbons
    const normMouseY = (mouseY / (height || 1)) - 0.5

    // 2. Render continuous flowing architectural contour isolines
    for (let i = 0; i < totalRibbons; i++) {
      // Calculate continuous looping baseY with fluid vertical flow
      const rawY = (i * spacing - continuousVerticalOffset) % totalSpan
      const baseY = (rawY < 0 ? rawY + totalSpan : rawY) - padding

      // Vertical edge fade to ensure wrapping ribbons emerge & vanish seamlessly
      const normalizedViewY = (baseY + padding) / totalSpan
      const edgeFade = Math.sin(Math.max(0, Math.min(1, normalizedViewY)) * Math.PI)

      if (edgeFade <= 0.02) continue // Skip offscreen lines

      const ribbonProgress = i / totalRibbons
      const isHeroSky = i % 6 === 2
      const isHeroIndigo = i % 6 === 5
      const isDashed = i % 4 === 1

      ctx.beginPath()

      if (isDashed) {
        ctx.setLineDash([6, 8])
      } else {
        ctx.setLineDash([])
      }

      // Sample points across horizontal width
      const step = 22
      const points = []

      for (let x = -step; x <= width + step; x += step) {
        const xNorm = x / width

        // Multi-harmonic flow equations
        const w1 = Math.sin(x * 0.0020 + time * 1.4 + ribbonProgress * 3.6) * (32 * dynamicAmp)
        const w2 = Math.cos(x * 0.0044 - time * 0.9 + ribbonProgress * 2.2) * (16 * dynamicAmp)
        const w3 = Math.sin(x * 0.0088 + time * 2.1 + i * 0.45) * (6 * dynamicAmp)

        // Fluid mouse repulsion
        const mDist = Math.hypot(x - mouseX, baseY - mouseY)
        const mInfluence = Math.max(0, 1 - mDist / 380) * 24 * -normMouseY

        // Natural geological ridge curvature
        const ridge = Math.sin(xNorm * Math.PI * 1.5 + (baseY / height) * 1.2) * 22

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

      // Dynamic opacity & styling
      if (isHeroSky) {
        const grad = ctx.createLinearGradient(0, 0, width, 0)
        const alpha = (isDarkMode ? 0.60 : 0.50) * edgeFade * (1.0 + normVel * 0.3)
        grad.addColorStop(0, `rgba(${colors.primary}, 0.0)`)
        grad.addColorStop(0.2, `rgba(${colors.primary}, ${alpha * 0.85})`)
        grad.addColorStop(0.5, `rgba(${colors.highlight}, ${alpha * 1.15})`)
        grad.addColorStop(0.8, `rgba(${colors.primary}, ${alpha * 0.85})`)
        grad.addColorStop(1, `rgba(${colors.primary}, 0.0)`)
        ctx.strokeStyle = grad
        ctx.lineWidth = isDarkMode ? 1.6 : 1.5
      } else if (isHeroIndigo) {
        const grad = ctx.createLinearGradient(0, 0, width, 0)
        const alpha = (isDarkMode ? 0.52 : 0.42) * edgeFade * (1.0 + normVel * 0.3)
        grad.addColorStop(0, `rgba(${colors.accent}, 0.0)`)
        grad.addColorStop(0.25, `rgba(${colors.accent}, ${alpha * 0.85})`)
        grad.addColorStop(0.65, `rgba(${colors.accent}, ${alpha * 1.1})`)
        grad.addColorStop(1, `rgba(${colors.accent}, 0.0)`)
        ctx.strokeStyle = grad
        ctx.lineWidth = isDarkMode ? 1.4 : 1.3
      } else if (isDashed) {
        ctx.strokeStyle = `rgba(${colors.neutral}, ${(isDarkMode ? 0.22 : 0.20) * edgeFade})`
        ctx.lineWidth = 1.0
      } else {
        const baseAlpha = (isDarkMode ? 0.16 : 0.15) * edgeFade
        const depthAlpha = baseAlpha + Math.sin(ribbonProgress * Math.PI) * 0.08 * edgeFade
        ctx.strokeStyle = `rgba(${colors.neutral}, ${depthAlpha})`
        ctx.lineWidth = 1.05
      }

      ctx.stroke()

      // Precision telemetry crosshairs (+) at harmonic peaks
      if (i % 5 === 2 && edgeFade > 0.4) {
        const checkIdx = Math.floor(points.length * (0.22 + (i % 3) * 0.26))
        if (points[checkIdx]) {
          const pt = points[checkIdx]
          const cSize = 4
          ctx.save()
          ctx.setLineDash([])
          ctx.strokeStyle = colors.crosshair
          ctx.lineWidth = 1.1
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

    ctx.setLineDash([])

    // 3. Dynamic traveling energy pulses
    energyPulses.forEach((pulse) => {
      const pulseSpeed = pulse.speed * speedBoost
      pulse.xProgress = (pulse.xProgress + (prefersReducedMotion ? 0 : pulseSpeed)) % 1.0

      const ribbonIdx = pulse.ribbonIndex
      const rawY = (ribbonIdx * spacing - continuousVerticalOffset) % totalSpan
      const baseY = (rawY < 0 ? rawY + totalSpan : rawY) - padding
      const normalizedViewY = (baseY + padding) / totalSpan
      const edgeFade = Math.sin(Math.max(0, Math.min(1, normalizedViewY)) * Math.PI)

      if (edgeFade > 0.15) {
        const centerX = pulse.xProgress * width
        const halfLen = pulse.length * 0.5
        const startX = Math.max(0, centerX - halfLen)
        const endX = Math.min(width, centerX + halfLen)

        if (startX < endX) {
          const pulseGrad = ctx.createLinearGradient(startX, 0, endX, 0)
          const pulseAlpha = (isDarkMode ? 0.75 : 0.65) * edgeFade * (1.0 + normVel * 0.25)
          pulseGrad.addColorStop(0, `rgba(${colors.primary}, 0)`)
          pulseGrad.addColorStop(0.5, `rgba(${colors.highlight}, ${pulseAlpha})`)
          pulseGrad.addColorStop(1, `rgba(${colors.primary}, 0)`)

          ctx.beginPath()
          for (let px = startX; px <= endX; px += 16) {
            const ribbonProgress = ribbonIdx / totalRibbons
            const w1 = Math.sin(px * 0.0020 + time * 1.4 + ribbonProgress * 3.6) * (32 * dynamicAmp)
            const w2 = Math.cos(px * 0.0044 - time * 0.9 + ribbonProgress * 2.2) * (16 * dynamicAmp)
            const w3 = Math.sin(px * 0.0088 + time * 2.1 + ribbonIdx * 0.45) * (6 * dynamicAmp)
            const py = baseY + w1 + w2 + w3

            if (px === startX) {
              ctx.moveTo(px, py)
            } else {
              ctx.lineTo(px, py)
            }
          }

          ctx.strokeStyle = pulseGrad
          ctx.lineWidth = isDarkMode ? 2.6 : 2.4
          ctx.stroke()
        }
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
  background: radial-gradient(circle at 50% 15%, #ffffff 0%, #f8fafc 45%, #e2e8f0 100%);
}

:root.dark .dynamic-background-container {
  background: radial-gradient(circle at 50% 12%, #0b1120 0%, #060911 55%, #020408 100%);
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
  background: radial-gradient(circle at center, transparent 30%, rgba(2, 4, 8, 0.75) 100%);
}

:root:not(.dark) .radial-vignette {
  background: radial-gradient(circle at center, transparent 40%, rgba(226, 232, 240, 0.55) 100%);
}
</style>
