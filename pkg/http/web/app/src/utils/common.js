export const SOCIAL_MEDIA_LIMIT = 10
export const SOCIAL_MEDIA_SKIP = 50

export const convertFromBinary = (type = "image/svg+xml", binary) => {
    return `data:${type};base64,${binary}`
}

export const chunk = (input, size) => {
    return input.reduce((arr, item, idx) => {
        return idx % size === 0 ?
            [...arr, [item]] :
            [...arr.slice(0, -1), [...arr.slice(-1)[0], item]];
    }, []);
};

export const shuffleArray = arr => {
    const newArr = arr.slice()
    for (let i = newArr.length - 1; i > 0; i--) {
        const rand = Math.floor(Math.random() * (i + 1));
        [newArr[i], newArr[rand]] = [newArr[rand], newArr[i]];
    }
    return newArr
};

export const openLink = (path) => {
    window.open(path, '_blank');
}