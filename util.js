const yaml = require('yaml')

function parse(resp){
    resp.results.forEach((event)=>{
        // print(yaml.stringify(event)+'\n') //print to go
        print(event.description+'\n') //print to go
        // return {
        //     id: event.id,
        //     slug: event.slug,
        //     eventtype: event.type.name,
        //     description: event.description,
        //     location: event.location,
        //     img: event.feature_image,
        //     date: event.date
        // }
        // return JSON.stringify(event.slug)
    })
}