import React from 'react'

interface LinkButtonProps {
    children: React.ReactNode
    className?: string
    onClick?: () => void
}

const LinkButton: React.FC<LinkButtonProps> = ({ children, className, onClick }) => {
    return (
        <button
            className={`font-nunito text-buttonDefault font-bold cursor-pointer hover:text-buttonHover ${
                className ? className : ''
            }`}
            onClick={onClick}
        >
            {children}
        </button>
    )
}

export default LinkButton
