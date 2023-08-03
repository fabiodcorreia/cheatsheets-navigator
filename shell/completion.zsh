function csn {
    go run main.go
}

function _csn_completion {
    if [[ -v CSN_PAGES ]]; then
        local pages=($(csn -ls))
        _describe 'values' pages
    fi
}

compdef _csn_completion csn
