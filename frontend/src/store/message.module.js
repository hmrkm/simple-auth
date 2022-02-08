const state = {
    content: '',
    clss: '',
}

const mutations = {
    setContent (state, { content, timeout, clss }) {
        state.content = content
        state.clss = clss

        if (typeof timeout === 'undefined') {
            timeout = 3000
        }

        setTimeout(() => {
            state.content = '';
            state.clss = '';
        }, timeout)
    }
}

export const message =  {
    namespaced: true,
    state,
    mutations
}