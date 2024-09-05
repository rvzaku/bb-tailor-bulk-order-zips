import React from 'react'

interface PrimaryButtonProps {
    children: React.ReactNode
    className?: string
    onClick?: () => void
}

const PrimaryButton: React.FC<PrimaryButtonProps> = ({ children, className, onClick }) => {
    return (
        <button
            className={`font-nunito text-buttonText pointer-cursor rounded-md bg-buttonDefault hover:bg-buttonHover ${className}`}
            onClick={onClick}
        >
            {children}
        </button>
    )
}

export default PrimaryButton
