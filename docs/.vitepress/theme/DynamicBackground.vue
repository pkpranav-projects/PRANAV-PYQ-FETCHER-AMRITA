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

  // Check dark mode
  const updateTheme = () => {
    isDarkMode = document.documentElement.classList.contains('dark')
  }
  updateTheme()

  const observer = new MutationObserver(() => updateTheme())
  observer.observe(document.documentElement, {
    attributes: true,
    attributeFilter: ['class'],
  })

  // Resize handler with devicePixelRatio support
  let width = 0
  let height = 0
  const handleResize = () => {
    const dpr = Math.min(window.devicePixelRatio || 1, 2)
    width = window.innerWidth
    height = window.innerHeight
    canvas.width = width * dpr
    canvas.height = height * dpr
    ctx.scale(dpr, dpr)
  }
  handleResize()
  window.addEventListener('resize', handleResize, { passive: true })

  // Smooth scroll interpolation
  let currentScroll = window.scrollY
  let targetScroll = window.scrollY
  const handleScroll = () => {
    targetScroll = window.scrollY
  }
  window.addEventListener('scroll', handleScroll, { passive: true })

  // Subtle mouse interaction
  let mouseX = width / 2
  let mouseY = height / 2
  let targetMouseX = width / 2
  let targetMouseY = height / 2
  const handleMouseMove = (e) => {
    targetMouseX = e.clientX
    targetMouseY = e.clientY
  }
  window.addEventListener('mousemove', handleMouseMove, { passive: true })

  // Check user preference for reduced motion
  const prefersReducedMotion = window.matchMedia('(prefers-reduced-motion: reduce)').matches

  // Topographic elevation centers (peaks)
  // These correspond to the concentric topographic circles in the user's reference image
  const peaks = [
    { xRatio: 0.18, yRatio: 0.15, radius: 140, rings: 7, pulseOffset: 0.0 },
    { xRatio: 0.45, yRatio: 0.28, radius: 120, rings: 6, pulseOffset: 1.5 },
    { xRatio: 0.82, yRatio: 0.18, radius: 150, rings: 8, pulseOffset: 3.0 },
    { xRatio: 0.22, yRatio: 0.72, radius: 130, rings: 6, pulseOffset: 4.5 },
    { xRatio: 0.78, yRatio: 0.78, radius: 160, rings: 8, pulseOffset: 2.2 },
  ]

  let time = 0

  const render = () => {
    time += prefersReducedMotion ? 0 : 0.006

    // Smooth lerp for scroll and mouse
    currentScroll += (targetScroll - currentScroll) * 0.08
    mouseX += (targetMouseX - mouseX) * 0.05
    mouseY += (targetMouseY - mouseY) * 0.05

    ctx.clearRect(0, 0, width, height)

    // Palette setup based on current theme
    const primaryStroke = isDarkMode ? 'rgba(1, 250, 198, ' : 'rgba(14, 165, 233, '
    const secondaryStroke = isDarkMode ? 'rgba(255, 255, 255, ' : 'rgba(15, 23, 42, '
    const accentStroke = isDarkMode ? 'rgba(168, 85, 247, ' : 'rgba(99, 102, 241, '

    // 1. Draw flowing horizontal contour waves
    // Number of horizontal contour lines across viewport
    const numLines = 26
    const lineSpacing = height / (numLines - 1)
    const scrollFactor = currentScroll * 0.0018

    ctx.lineWidth = 1.1

    for (let i = 0; i < numLines; i++) {
      const baseY = i * lineSpacing
      ctx.beginPath()

      // Cycle colors: occasionally highlight with brand cyan or violet
      let alpha = 0.09
      let strokeColor = secondaryStroke + alpha + ')'

      if (i % 5 === 0) {
        alpha = isDarkMode ? 0.25 : 0.22
        strokeColor = primaryStroke + alpha + ')'
      } else if (i % 7 === 0) {
        alpha = isDarkMode ? 0.18 : 0.16
        strokeColor = accentStroke + alpha + ')'
      }

      ctx.strokeStyle = strokeColor

      // Step along the horizontal width
      const step = 20
      for (let x = 0; x <= width + step; x += step) {
        // Base sine wave modulation influenced by scroll
        const wave1 = Math.sin(x * 0.003 + scrollFactor + time + i * 0.2) * 14
        const wave2 = Math.cos(x * 0.006 - scrollFactor * 0.7 + i * 0.3) * 8

        // Elevation displacement from each topographic peak
        let peakElevation = 0
        peaks.forEach((peak) => {
          const px = peak.xRatio * width
          const py = peak.yRatio * height + Math.sin(scrollFactor + peak.pulseOffset) * 20
          const dist = Math.hypot(x - px, baseY - py)
          if (dist < peak.radius * 1.8) {
            const influence = (1 - dist / (peak.radius * 1.8))
            peakElevation += Math.sin(dist * 0.12 - time * 2) * influence * 12
          }
        })

        const y = baseY + wave1 + wave2 + peakElevation

        if (x === 0) {
          ctx.moveTo(x, y)
        } else {
          ctx.lineTo(x, y)
        }
      }
      ctx.stroke()
    }

    // 2. Draw topographic concentric elevation rings (The islands from reference image)
    peaks.forEach((peak, pIndex) => {
      const px = peak.xRatio * width + (mouseX - width / 2) * 0.02
      // Parallax scroll on peaks
      const scrollShift = (currentScroll * (0.15 + (pIndex % 3) * 0.05)) % (height * 0.4)
      const py = peak.yRatio * height - scrollShift + Math.sin(time + peak.pulseOffset) * 8

      // Draw concentric rings for this elevation peak
      for (let r = 1; r <= peak.rings; r++) {
        const ringProgress = r / peak.rings
        const baseRadius = ringProgress * peak.radius
        // Morph ring with scroll and time
        const dynamicRadius = baseRadius + Math.sin(scrollFactor * 3 + r * 0.8 + time) * 6

        ctx.beginPath()

        // Distort circular shape into natural organic topographic contour
        const numPoints = 48
        for (let pt = 0; pt <= numPoints; pt++) {
          const angle = (pt / numPoints) * Math.PI * 2
          // Organic wobble factor
          const wobble =
            Math.sin(angle * 3 + time + r) * 7 +
            Math.cos(angle * 2 - scrollFactor * 2) * 5

          const currentR = Math.max(8, dynamicRadius + wobble)
          const rx = px + Math.cos(angle) * currentR * 1.35 // Slightly stretched ellipse
          const ry = py + Math.sin(angle) * currentR

          if (pt === 0) {
            ctx.moveTo(rx, ry)
          } else {
            ctx.lineTo(rx, ry)
          }
        }
        ctx.closePath()

        let ringAlpha = 0.10 + (r / peak.rings) * 0.16
        if (r === peak.rings || r === 1) {
          ctx.strokeStyle = primaryStroke + (ringAlpha * 1.4) + ')'
          ctx.lineWidth = 1.3
        } else if (r % 2 === 0) {
          ctx.strokeStyle = accentStroke + (ringAlpha * 1.1) + ')'
          ctx.lineWidth = 1.0
        } else {
          ctx.strokeStyle = secondaryStroke + ringAlpha + ')'
          ctx.lineWidth = 0.9
        }

        ctx.stroke()
      }
    })

    animationFrameId = requestAnimationFrame(render)
  }

  render()

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
  background: radial-gradient(circle at 50% 20%, #ffffff 0%, #f1f5f9 60%, #e2e8f0 100%);
}

:root.dark .dynamic-background-container {
  background: radial-gradient(circle at 50% 15%, #0d1322 0%, #080c14 60%, #030508 100%);
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
  background: radial-gradient(circle at center, transparent 30%, rgba(3, 5, 8, 0.65) 100%);
}

:root:not(.dark) .radial-vignette {
  background: radial-gradient(circle at center, transparent 40%, rgba(226, 232, 240, 0.45) 100%);
}
</style>
