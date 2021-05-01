// const yaml = require('yaml')

function parse(resp){
    print(toYaml(resp,0))
    // resp.results.forEach((event)=>{
    //     // print(JSON.stringify(event)+'\n') 

    //     // print(event.description+'\n') 
    //     // return {
    //     //     id: event.id,
    //     //     slug: event.slug,
    //     //     eventtype: event.type.name,
    //     //     description: event.description,
    //     //     location: event.location,
    //     //     img: event.feature_image,
    //     //     date: event.date
    //     // }
    //     // return JSON.stringify(event.slug)
    // })
}

function toYaml(data,depth){
    console.log(`depth: $(depth)`)
    const TAB = '   '
    res = ''
    keys = Object.keys(data)
    keys.forEach((key)=>{
        var item = data[key]
        if (item === null){
            res+=`${TAB.repeat(depth)}${key}: {}\n`
        } else if (typeof(item) == 'string'){
            res+=`${TAB.repeat(depth)}${key}: "${item}"\n`
        }else if (typeof(item) == 'number'){
            res+=`${TAB.repeat(depth)}${key}: ${item}\n`
        } else {
            res+=toYaml(item,depth+1)
        }
    })
    // console.log(res)
    return res
}