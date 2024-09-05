import { useState, useEffect } from 'react'

const useScreenSize = (defaultWidth: string, defaultHeight: string) => {
    const [size, setSize] = useState<{ width: string; height: string }>({
        width: defaultWidth,
        height: defaultHeight,
    })

    useEffect(() => {
        const handleResize = () => {
            const width = window.innerWidth
            let newWidth = defaultWidth
            let newHeight = defaultHeight

            if (width >= 640) {
                // sm breakpoint
                newWidth = '75px'
                newHeight = '75px'
            }
            if (width >= 768) {
                // md breakpoint
                newWidth = '100px'
                newHeight = '100px'
            }
            if (width >= 1024) {
                // lg breakpoint
                newWidth = '100px'
                newHeight = '100px'
            }
            if (width >= 1280) {
                // xl breakpoint
                newWidth = '150px'
                newHeight = '150px'
            }
            if (width >= 1536) {
                // 2xl breakpoint
                newWidth = '150px'
                newHeight = '150px'
            }

            setSize({ width: newWidth, height: newHeight })
        }

        handleResize() // Set the initial size
        window.addEventListener('resize', handleResize)
        return () => window.removeEventListener('resize', handleResize)
    }, [defaultWidth, defaultHeight])

    return size
}

export default useScreenSize
