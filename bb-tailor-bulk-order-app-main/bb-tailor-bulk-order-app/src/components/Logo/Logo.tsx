import React from 'react'

interface LogoProps {
    className?: string
    width?: string
    height?: string
}

const Logo: React.FC<LogoProps> = ({ className, width, height }) => {
    return (
        <img
            className={`${className}`}
            alt="Brooks logo"
            src="../../../public/brooks_logo.svg"
            height={height}
            width={width}
        />
    )
}

export default Logo
