import React from 'react'

interface InputProps {
    className?: string
    label?: string
    type?: string
    name?: string
    placeholder?: string
    value?: string
    onChange?: (event: React.ChangeEvent<HTMLInputElement>) => void
}

const Input: React.FC<InputProps> = ({ className, label, type, name, placeholder, value, onChange }) => {
    return (
        <div className="mb-4">
            {label && <label className="font-nunito block text-sm font-medium mb-2">{label}</label>}
            <input
                className={`font-nunito w-full bg-primaryBrand border border-secondaryText rounded px-3 py-2 text-secondaryText focus:outline-none focus:ring-1 ${
                    className ? className : ''
                }`}
                type={type || 'text'}
                name={name}
                placeholder={placeholder}
                value={value}
                onChange={onChange}
            />
        </div>
    )
}

export default Input
