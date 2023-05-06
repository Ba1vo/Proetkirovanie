export const objectToQueryString = (object) => {
  const queries = []
  for (const key in object) {
    if (object[key] && object[key]?.length) {
      let query = `${key}=`
      if (Array.isArray(object[key])) {
        if (key === 'volume') {
          object[key] = object[key].map((value) => Number(value) * 1000)
        }
        query += object[key].join(';')
      } else {
        query += object[key]
      }
      queries.push(query)
    }
  }
  return queries.join('&')
}

export const queryStringToObject = (queryString) => {
  const object = Object.fromEntries(new URLSearchParams(queryString))
  for (const key in object) {
    if (object[key].includes(';')) {
      object[key] = object[key].split(';')
    }
  }
  return object
}
