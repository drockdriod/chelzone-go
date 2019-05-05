
export const convertFromBinary = (type="image/svg+xml", binary) => {
    return `data:${type};base64,${binary}`
}