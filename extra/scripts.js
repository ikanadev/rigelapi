// those are the scripts used to get school data
// from the site: https://seie.minedu.gob.bo/reportes/mapas_unidades_educativas/#
// enter to the site console and run these scripts
await fetch("http://localhost:4000/deps", {
  method: 'POST',
  body: JSON.stringify(GeoArr.dep.map(d => ({ id: `${d.cdep}`, name: d.text }))),
  headers: {
    'Access-Control-Allow-Origin': '*',
    'Content-Type': 'application/json'
  }
})

await fetch("http://localhost:4000/provs", {
  method: 'POST',
  body: JSON.stringify(GeoArr.pro.map(p => ({ name: p.text, dpto_id: `${p.cdep}`, id: `${p.cid}` }))),
  headers: {
    'Access-Control-Allow-Origin': '*',
    'Content-Type': 'application/json'
  }
})

await fetch("http://localhost:4000/muns", {
  method: 'POST',
  body: JSON.stringify(GeoArr.mun.map(m => ({ name: m.text, prov_id: `${m.fid}`, id: `${m.cmun}` }))),
  headers: {
    'Access-Control-Allow-Origin': '*',
    'Content-Type': 'application/json'
  }
})

await fetch("http://localhost:4000/cols", {
  method: 'POST',
  body: JSON.stringify(GeoArr.unidad_geo.map(c => ({ id: c.cod_ue, lat: c.latitud, lon: c.longitud, name: c.name, mun_id: `${c.fidmun}` }))),
  headers: {
    'Access-Control-Allow-Origin': '*',
    'Content-Type': 'application/json'
  }
})
