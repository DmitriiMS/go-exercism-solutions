package letter

import (
	"reflect"
	"testing"
)

// In the separate file parallel_letter_frequency.go, you are given a function,
// Frequency(), to sequentially count letter frequencies in a single text.
//
// Perform this exercise on parallelism using Go concurrency features.
// Make concurrent calls to Frequency and combine results to obtain the answer.

var (
	euro = `

	Wiki Loves Monuments: your chance to support Russian cultural heritage! Photograph a monument and win!
	Featured article
	Page semi-protected
	Listen to this article
	Mercury (planet)
	From Wikipedia, the free encyclopedia
	Jump to navigation
	Jump to search
	Mercury Astronomical symbol of MercuryMercury in true color.jpg
	Mercury in true color (by MESSENGER in 2008)
	Designations
	Pronunciation	/ˈmɜːrkjʊri/ (About this soundlisten)
	Adjectives	Mercurian /mərˈkjʊəriən/,[1]
	Mercurial /mərˈkjʊəriəl/[2]
	Orbital characteristics[5]
	Epoch J2000
	Aphelion	
	
		0.466697 AU
		69,816,900 km
	
	Perihelion	
	
		0.307499 AU
		46,001,200 km
	
	Semi-major axis
		
	
		0.387098 AU
		57,909,050 km
	
	Eccentricity	0.205630[3]
	Orbital period
		
	
		87.9691 d
		0.240846 yr
		0.5 Mercury synodic day
	
	Synodic period
		115.88 d[3]
	Average orbital speed
		47.36 km/s[3]
	Mean anomaly
		174.796°
	Inclination	
	
		7.005° to ecliptic
		3.38° to Sun's equator
		6.35° to invariable plane[4]
	
	Longitude of ascending node
		48.331°
	Argument of perihelion
		29.124°
	Satellites	None
	Physical characteristics
	Mean diameter
		4880 km
	Mean radius
		
	
		2,439.7±1.0 km[6][7]
		0.3829 Earths
	
	Flattening	0.0000[3]
	Surface area
		
	
		7.48×107 km2[6]
		0.147 Earths
	
	Volume	
	
		6.083×1010 km3[6]
		0.056 Earths
	
	Mass	
	
		3.3011×1023 kg[8]
		0.055 Earths
	
	Mean density
		5.427 g/cm3[6]
	Surface gravity
		
	
		3.7 m/s2
		0.38 g[6]
	
	Moment of inertia factor
		0.346±0.014[9]
	Escape velocity
		4.25 km/s[6]
	Rotation period
		176 d (synodic; solar day)[10]
	Sidereal rotation period
		
	
		58.646 d
		1407.5 h[6]
	
	Equatorial rotation velocity
		10.892 km/h (3.026 m/s)
	Axial tilt
		2.04′ ± 0.08′ (to orbit)[9]
	(0.034°)[3]
	North pole right ascension
		
	
		18h 44m 2s
		281.01°[3]
	
	North pole declination
		61.45°[3]
	Albedo	
	
		0.088 (Bond)[11]
		0.142 (geom.)[12]
	
	Surface temp. 	min 	mean 	max
	0°N, 0°W [13] 	-173 °C 	67 °C 	427 °C
	85°N, 0°W[13] 	-193 °C 	-73 °C 	106.85 °C
	Apparent magnitude
		−2.48 to +7.25[14]
	Angular diameter
		4.5–13″[3]
	Atmosphere[3][15][16]
	Surface pressure
		trace (≲ 0.5 nPa)
	Composition by volume	
	
		atomic oxygen
		sodium
		magnesium
		atomic hydrogen
		potassium
		calcium
		helium
		Trace amounts of iron, aluminium, argon, dinitrogen, dioxygen, carbon dioxide, water vapor, xenon, krypton, and neon
	
	Mercury is the smallest planet in the Solar System and the closest to the Sun. Its orbit around the Sun takes 87.97 Earth days, the shortest of all the Sun's planets. It is named after the Roman god Mercurius (Mercury), god of commerce, messenger of the gods, and mediator between gods and mortals, corresponding to the Greek god Hermes (Ἑρμῆς). Like Venus, Mercury orbits the Sun within Earth's orbit as an inferior planet, and its apparent distance from the Sun as viewed from Earth never exceeds 28°. This proximity to the Sun means the planet can only be seen near the western horizon after sunset or the eastern horizon before sunrise, usually in twilight. At this time, it may appear as a bright star-like object but is often far more difficult to observe than Venus. From Earth, the planet telescopically displays the complete range of phases, similar to Venus and the Moon, which recurs over its synodic period of approximately 116 days.
	
	Mercury rotates in a way that is unique in the Solar System. It is tidally locked with the Sun in a 3:2 spin–orbit resonance,[17] meaning that relative to the fixed stars, it rotates on its axis exactly three times for every two revolutions it makes around the Sun.[a][18] As seen from the Sun, in a frame of reference that rotates with the orbital motion, it appears to rotate only once every two Mercurian years. An observer on Mercury would therefore see only one day every two Mercurian years.
	
	Mercury's axis has the smallest tilt of any of the Solar System's planets (about 1⁄30 degree). Its orbital eccentricity is the largest of all known planets in the Solar System;[b] at perihelion, Mercury's distance from the Sun is only about two-thirds (or 66%) of its distance at aphelion. Mercury's surface appears heavily cratered and is similar in appearance to the Moon's, indicating that it has been geologically inactive for billions of years. Having almost no atmosphere to retain heat, it has surface temperatures that vary diurnally more than on any other planet in the Solar System, ranging from 100 K (−173 °C; −280 °F) at night to 700 K (427 °C; 800 °F) during the day across the equatorial regions.[19] The polar regions are constantly below 180 K (−93 °C; −136 °F). The planet has no known natural satellites.
	
	Two spacecraft have visited Mercury: Mariner 10 flew by in 1974 and 1975; and MESSENGER, launched in 2004, orbited Mercury over 4,000 times in four years before exhausting its fuel and crashing into the planet's surface on April 30, 2015.[20][21][22] The BepiColombo spacecraft is planned to arrive at Mercury in 2025.
	Contents
	
		1 Physical characteristics
			1.1 Internal structure
			1.2 Surface geology
				1.2.1 Impact basins and craters
				1.2.2 Plains
				1.2.3 Compressional features
				1.2.4 Volcanism
			1.3 Surface conditions and exosphere
			1.4 Magnetic field and magnetosphere
		2 Orbit, rotation, and longitude
			2.1 Longitude convention
			2.2 Spin-orbit resonance
			2.3 Advance of perihelion
		3 Habitability
		4 Observation
		5 Observation history
			5.1 Ancient astronomers
			5.2 Ground-based telescopic research
			5.3 Research with space probes
				5.3.1 Mariner 10
				5.3.2 MESSENGER
				5.3.3 BepiColombo
		6 Comparison
		7 See also
		8 Notes
		9 References
		10 External links
	
	Physical characteristics
	
	Mercury is one of four terrestrial planets in the Solar System, and is a rocky body like Earth. It is the smallest planet in the Solar System, with an equatorial radius of 2,439.7 kilometres (1,516.0 mi).[3] Mercury is also smaller—albeit more massive—than the largest natural satellites in the Solar System, Ganymede and Titan. Mercury consists of approximately 70% metallic and 30% silicate material.[23]
	Internal structure
	Mercury's internal structure and magnetic field
	
	Mercury appears to have a solid silicate crust and mantle overlying a solid, iron sulfide outer core layer, a deeper liquid core layer, and a solid inner core.[24][25] The planet's density is the second highest in the Solar System at 5.427 g/cm3, only slightly less than Earth's density of 5.515 g/cm3.[3] If the effect of gravitational compression were to be factored out from both planets, the materials of which Mercury is made would be denser than those of Earth, with an uncompressed density of 5.3 g/cm3 versus Earth's 4.4 g/cm3.[26] Mercury's density can be used to infer details of its inner structure. Although Earth's high density results appreciably from gravitational compression, particularly at the core, Mercury is much smaller and its inner regions are not as compressed. Therefore, for it to have such a high density, its core must be large and rich in iron.[27]
	
	Geologists estimate that Mercury's core occupies about 55% of its volume; for Earth this proportion is 17%. Research published in 2007 suggests that Mercury has a molten core.[28][29] Surrounding the core is a 500–700 km (310–430 mi) mantle consisting of silicates.[30][31] Based on data from the Mariner 10 mission and Earth-based observation, Mercury's crust is estimated to be 35 km (22 mi) thick.[32] However, this model may be an overestimate and the crust could be 26 ± 11 km (16.2 ± 6.8 mi) thick based on an Airy isostacy model.[33] One distinctive feature of Mercury's surface is the presence of numerous narrow ridges, extending up to several hundred kilometers in length. It is thought that these were formed as Mercury's core and mantle cooled and contracted at a time when the crust had already solidified.[34][35][36]
	
	Mercury's core has a higher iron content than that of any other major planet in the Solar System, and several theories have been proposed to explain this. The most widely accepted theory is that Mercury originally had a metal–silicate ratio similar to common chondrite meteorites, thought to be typical of the Solar System's rocky matter, and a mass approximately 2.25 times its current mass.[37] Early in the Solar System's history, Mercury may have been struck by a planetesimal of approximately 1/6 that mass and several thousand kilometers across.[37] The impact would have stripped away much of the original crust and mantle, leaving the core behind as a relatively major component.[37] A similar process, known as the giant impact hypothesis, has been proposed to explain the formation of the Moon.[37]
	
	Alternatively, Mercury may have formed from the solar nebula before the Sun's energy output had stabilized. It would initially have had twice its present mass, but as the protosun contracted, temperatures near Mercury could have been between 2,500 and 3,500 K and possibly even as high as 10,000 K.[38] Much of Mercury's surface rock could have been vaporized at such temperatures, forming an atmosphere of "rock vapor" that could have been carried away by the solar wind.[38]
	
	A third hypothesis proposes that the solar nebula caused drag on the particles from which Mercury was accreting, which meant that lighter particles were lost from the accreting material and not gathered by Mercury.[39] Each hypothesis predicts a different surface composition, and there are two space missions set to make observations. MESSENGER, which ended in 2015, found higher-than-expected potassium and sulfur levels on the surface, suggesting that the giant impact hypothesis and vaporization of the crust and mantle did not occur because potassium and sulfur would have been driven off by the extreme heat of these events.[40] BepiColombo, which will arrive at Mercury in 2025, will make observations to test these hypotheses.[41] The findings so far would seem to favor the third hypothesis; however, further analysis of the data is needed.[42]
	Surface geology
	Main article: Geology of Mercury
	
	Mercury's surface is similar in appearance to that of the Moon, showing extensive mare-like plains and heavy cratering, indicating that it has been geologically inactive for billions of years. It is more heterogeneous than either Mars's or the Moon's, both of which contain significant stretches of similar geology, such as maria and plateaus.[43] Albedo features are areas of markedly different reflectivity, which include impact craters, the resulting ejecta, and ray systems. Larger albedo features correspond to higher reflectivity plains.[44] Mercury has dorsa (also called "wrinkle-ridges"), Moon-like highlands, montes (mountains), planitiae (plains), rupes (escarpments), and valles (valleys).[45][46]
	MASCS spectrum scan of Mercury's surface by MESSENGER
	
	The planet's mantle is chemically heterogeneous, suggesting the planet went through a magma ocean phase early in its history. Crystallization of minerals and convective overturn resulted in layered, chemically heterogeneous crust with large-scale variations in chemical composition observed on the surface. The crust is low in iron but high in sulfur, resulting from the stronger early chemically reducing conditions than is found in the other terrestrial planets. The surface is dominated by iron-poor pyroxene and olivine, as represented by enstatite and forsterite, respectively, along with sodium-rich plagioclase and minerals of mixed magnesium, calcium, and iron-sulfide. The less reflective regions of the crust are high in carbon, most likely in the form of graphite.[47]
	
	Names for features on Mercury come from a variety of sources. Names coming from people are limited to the deceased. Craters are named for artists, musicians, painters, and authors who have made outstanding or fundamental contributions to their field. Ridges, or dorsa, are named for scientists who have contributed to the study of Mercury. Depressions or fossae are named for works of architecture. Montes are named for the word "hot" in a variety of languages. Plains or planitiae are named for Mercury in various languages. Escarpments or rupēs are named for ships of scientific expeditions. Valleys or valles are named for abandoned cities, towns, or settlements of antiquity.[48]
	Impact basins and craters
	Enhanced-color image of Munch, Sander and Poe craters amid volcanic plains (orange) near Caloris Basin
	
	Mercury was heavily bombarded by comets and asteroids during and shortly following its formation 4.6 billion years ago, as well as during a possibly separate subsequent episode called the Late Heavy Bombardment that ended 3.8 billion years ago.[49] Mercury received impacts over its entire surface during this period of intense crater formation,[46] facilitated by the lack of any atmosphere to slow impactors down.[50] During this time Mercury was volcanically active; basins were filled by magma, producing smooth plains similar to the maria found on the Moon.[51][52] An unusual crater with radiating troughs has been discovered that scientists called "the spider".[53] It was later named Apollodorus.[54]
	
	Craters on Mercury range in diameter from small bowl-shaped cavities to multi-ringed impact basins hundreds of kilometers across. They appear in all states of degradation, from relatively fresh rayed craters to highly degraded crater remnants. Mercurian craters differ subtly from lunar craters in that the area blanketed by their ejecta is much smaller, a consequence of Mercury's stronger surface gravity.[55] According to International Astronomical Union (IAU) rules, each new crater must be named after an artist who was famous for more than fifty years, and dead for more than three years, before the date the crater is named.[56]
	Overhead view of Caloris Basin
	Perspective view of Caloris Basin – high (red); low (blue)
	
	The largest known crater is Caloris Planitia, or Caloris Basin, with a diameter of 1,550 km.[57] The impact that created the Caloris Basin was so powerful that it caused lava eruptions and left a concentric mountainous ring ~2 km tall surrounding the impact crater. The floor of the Caloris Basin is filled by a geologically distinct flat plain, broken up by ridges and fractures in a roughly polygonal pattern. It is not clear whether they are volcanic lava flows induced by the impact or a large sheet of impact melt.[55]
	
	At the antipode of the Caloris Basin is a large region of unusual, hilly terrain known as the "Weird Terrain". One hypothesis for its origin is that shock waves generated during the Caloris impact traveled around Mercury, converging at the basin's antipode (180 degrees away). The resulting high stresses fractured the surface.[58] Alternatively, it has been suggested that this terrain formed as a result of the convergence of ejecta at this basin's antipode.[59]
	Tolstoj basin is along the bottom of this image of Mercury's limb
	
	Overall, 46 impact basins have been identified.[60] A notable basin is the 400 km wide, multi-ring Tolstoj Basin that has an ejecta blanket extending up to 500 km from its rim and a floor that has been filled by smooth plains materials. Beethoven Basin has a similar-sized ejecta blanket and a 625 km diameter rim.[55] Like the Moon, the surface of Mercury has likely incurred the effects of space weathering processes, including solar wind and micrometeorite impacts.[61]
	Plains
	
	There are two geologically distinct plains regions on Mercury.[55][62] Gently rolling, hilly plains in the regions between craters are Mercury's oldest visible surfaces,[55] predating the heavily cratered terrain. These inter-crater plains appear to have obliterated many earlier craters, and show a general paucity of smaller craters below about 30 km in diameter.[62]
	
	Smooth plains are widespread flat areas that fill depressions of various sizes and bear a strong resemblance to the lunar maria. Unlike lunar maria, the smooth plains of Mercury have the same albedo as the older inter-crater plains. Despite a lack of unequivocally volcanic characteristics, the localisation and rounded, lobate shape of these plains strongly support volcanic origins.[55] All the smooth plains of Mercury formed significantly later than the Caloris basin, as evidenced by appreciably smaller crater densities than on the Caloris ejecta blanket.[55]
	Compressional features
	
	One unusual feature of Mercury's surface is the numerous compression folds, or rupes, that crisscross the plains. As Mercury's interior cooled, it contracted and its surface began to deform, creating wrinkle ridges and lobate scarps associated with thrust faults. The scarps can reach lengths of 1000 km and heights of 3 km.[63] These compressional features can be seen on top of other features, such as craters and smooth plains, indicating they are more recent.[64] Mapping of the features has suggested a total shrinkage of Mercury's radius in the range of ~1 to 7 km.[65] Most activity along the major thrust systems probably ended about 3.6–3.7 billion years ago.[66] Small-scale thrust fault scarps have been found, tens of meters in height and with lengths in the range of a few km, that appear to be less than 50 million years old, indicating that compression of the interior and consequent surface geological activity continue to the present.[63][65]
	
	The Lunar Reconnaissance Orbiter discovered that similar but smaller thrust faults exist on the Moon.[67]
	Volcanism
	Picasso crater — the large arc-shaped pit located on the eastern side of its floor are postulated to have formed when subsurface magma subsided or drained, causing the surface to collapse into the resulting void.
	
	There is evidence for pyroclastic flows on Mercury from low-profile shield volcanoes.[68][69][70] 51 pyroclastic deposits have been identified,[71] where 90% of them are found within impact craters.[71] A study of the degradation state of the impact craters that host pyroclastic deposits suggests that pyroclastic activity occurred on Mercury over a prolonged interval.[71]
	
	A "rimless depression" inside the southwest rim of the Caloris Basin consists of at least nine overlapping volcanic vents, each individually up to 8 km in diameter. It is thus a "compound volcano".[72] The vent floors are at least 1 km below their brinks and they bear a closer resemblance to volcanic craters sculpted by explosive eruptions or modified by collapse into void spaces created by magma withdrawal back down into a conduit.[72] Scientists could not quantify the age of the volcanic complex system but reported that it could be of the order of a billion years.[72]
	Surface conditions and exosphere
	Main article: Atmosphere of Mercury
	Composite of the north pole of Mercury, where NASA confirmed the discovery of a large volume of water ice, in permanently dark craters that are found there.[73]
	
	The surface temperature of Mercury ranges from 100 to 700 K (−173 to 427 °C; −280 to 800 °F)[19] at the most extreme places: 0°N, 0°W, or 180°W. It never rises above 180 K at the poles,[13] due to the absence of an atmosphere and a steep temperature gradient between the equator and the poles. The subsolar point reaches about 700 K during perihelion (0°W or 180°W), but only 550 K at aphelion (90° or 270°W).[74] On the dark side of the planet, temperatures average 110 K.[13][75] The intensity of sunlight on Mercury's surface ranges between 4.59 and 10.61 times the solar constant (1,370 W·m−2).[76]
	
	Although the daylight temperature at the surface of Mercury is generally extremely high, observations strongly suggest that ice (frozen water) exists on Mercury. The floors of deep craters at the poles are never exposed to direct sunlight, and temperatures there remain below 102 K, far lower than the global average.[77] This creates a cold trap where ice can accumulate. Water ice strongly reflects radar, and observations by the 70-meter Goldstone Solar System Radar and the VLA in the early 1990s revealed that there are patches of high radar reflection near the poles.[78] Although ice was not the only possible cause of these reflective regions, astronomers think it was the most likely.[79]
	
	The icy regions are estimated to contain about 1014–1015 kg of ice,[80] and may be covered by a layer of regolith that inhibits sublimation.[81] By comparison, the Antarctic ice sheet on Earth has a mass of about 4×1018 kg, and Mars's south polar cap contains about 1016 kg of water.[80] The origin of the ice on Mercury is not yet known, but the two most likely sources are from outgassing of water from the planet's interior or deposition by impacts of comets.[80]
	
	Mercury is too small and hot for its gravity to retain any significant atmosphere over long periods of time; it does have a tenuous surface-bounded exosphere[82] containing hydrogen, helium, oxygen, sodium, calcium, potassium and others[15][16] at a surface pressure of less than approximately 0.5 nPa (0.005 picobars).[3] This exosphere is not stable—atoms are continuously lost and replenished from a variety of sources. Hydrogen atoms and helium atoms probably come from the solar wind, diffusing into Mercury's magnetosphere before later escaping back into space. Radioactive decay of elements within Mercury's crust is another source of helium, as well as sodium and potassium. MESSENGER found high proportions of calcium, helium, hydroxide, magnesium, oxygen, potassium, silicon and sodium. Water vapor is present, released by a combination of processes such as: comets striking its surface, sputtering creating water out of hydrogen from the solar wind and oxygen from rock, and sublimation from reservoirs of water ice in the permanently shadowed polar craters. The detection of high amounts of water-related ions like O+, OH−, and H3O+ was a surprise.[83][84] Because of the quantities of these ions that were detected in Mercury's space environment, scientists surmise that these molecules were blasted from the surface or exosphere by the solar wind.[85][86]
	
	Sodium, potassium and calcium were discovered in the atmosphere during the 1980–1990s, and are thought to result primarily from the vaporization of surface rock struck by micrometeorite impacts[87] including presently from Comet Encke.[88] In 2008, magnesium was discovered by MESSENGER.[89] Studies indicate that, at times, sodium emissions are localized at points that correspond to the planet's magnetic poles. This would indicate an interaction between the magnetosphere and the planet's surface.[90]
	
	On November 29, 2012, NASA confirmed that images from MESSENGER had detected that craters at the north pole contained water ice. MESSENGER's principal investigator Sean Solomon is quoted in The New York Times estimating the volume of the ice to be large enough to "encase Washington, D.C., in a frozen block two and a half miles deep".[73]
	Magnetic field and magnetosphere
	Main article: Mercury's magnetic field
	Graph showing relative strength of Mercury's magnetic field
	
	Despite its small size and slow 59-day-long rotation, Mercury has a significant, and apparently global, magnetic field. According to measurements taken by Mariner 10, it is about 1.1% the strength of Earth's. The magnetic-field strength at Mercury's equator is about 300 nT.[91][92] Like that of Earth, Mercury's magnetic field is dipolar.[90] Unlike Earth's, Mercury's poles are nearly aligned with the planet's spin axis.[93] Measurements from both the Mariner 10 and MESSENGER space probes have indicated that the strength and shape of the magnetic field are stable.[93]
	
	It is likely that this magnetic field is generated by a dynamo effect, in a manner similar to the magnetic field of Earth.[94][95] This dynamo effect would result from the circulation of the planet's iron-rich liquid core. Particularly strong tidal heating effects caused by the planet's high orbital eccentricity would serve to keep part of the core in the liquid state necessary for this dynamo effect.[30][96]
	
	Mercury's magnetic field is strong enough to deflect the solar wind around the planet, creating a magnetosphere. The planet's magnetosphere, though small enough to fit within Earth,[90] is strong enough to trap solar wind plasma. This contributes to the space weathering of the planet's surface.[93] Observations taken by the Mariner 10 spacecraft detected this low energy plasma in the magnetosphere of the planet's nightside. Bursts of energetic particles in the planet's magnetotail indicate a dynamic quality to the planet's magnetosphere.[90]
	
	During its second flyby of the planet on October 6, 2008, MESSENGER discovered that Mercury's magnetic field can be extremely "leaky". The spacecraft encountered magnetic "tornadoes" – twisted bundles of magnetic fields connecting the planetary magnetic field to interplanetary space – that were up to 800 km wide or a third of the radius of the planet. These twisted magnetic flux tubes, technically known as flux transfer events, form open windows in the planet's magnetic shield through which the solar wind may enter and directly impact Mercury's surface via magnetic reconnection[97] This also occurs in Earth's magnetic field. The MESSENGER observations showed the reconnection rate is ten times higher at Mercury, but its proximity to the Sun only accounts for about a third of the reconnection rate observed by MESSENGER.[97]
	Orbit, rotation, and longitude
	Orbit of Mercury (2006)
	Animation of Mercury's and Earth's revolution around the Sun
	
	Mercury has the most eccentric orbit of all the planets in the Solar System; its eccentricity is 0.21 with its distance from the Sun ranging from 46,000,000 to 70,000,000 km (29,000,000 to 43,000,000 mi). It takes 87.969 Earth days to complete an orbit. The diagram illustrates the effects of the eccentricity, showing Mercury's orbit overlaid with a circular orbit having the same semi-major axis. Mercury's higher velocity when it is near perihelion is clear from the greater distance it covers in each 5-day interval. In the diagram, the varying distance of Mercury to the Sun is represented by the size of the planet, which is inversely proportional to Mercury's distance from the Sun. This varying distance to the Sun leads to Mercury's surface being flexed by tidal bulges raised by the Sun that are about 17 times stronger than the Moon's on Earth.[98] Combined with a 3:2 spin–orbit resonance of the planet's rotation around its axis, it also results in complex variations of the surface temperature.[23] The resonance makes a single solar day (the length between two meridian transits of the Sun) on Mercury last exactly two Mercury years, or about 176 Earth days.[99]
	
	Mercury's orbit is inclined by 7 degrees to the plane of Earth's orbit (the ecliptic), the largest of all eight known solar planets.[100] As a result, transits of Mercury across the face of the Sun can only occur when the planet is crossing the plane of the ecliptic at the time it lies between Earth and the Sun, which is in May or November. This occurs about every seven years on average.[101]
	
	Mercury's axial tilt is almost zero,[102] with the best measured value as low as 0.027 degrees.[103] This is significantly smaller than that of Jupiter, which has the second smallest axial tilt of all planets at 3.1 degrees. This means that to an observer at Mercury's poles, the center of the Sun never rises more than 2.1 arcminutes above the horizon.[103]
	
	At certain points on Mercury's surface, an observer would be able to see the Sun peek up a little more than two-thirds of the way over the horizon, then reverse and set before rising again, all within the same Mercurian day.[c] This is because approximately four Earth days before perihelion, Mercury's angular orbital velocity equals its angular rotational velocity so that the Sun's apparent motion ceases; closer to perihelion, Mercury's angular orbital velocity then exceeds the angular rotational velocity. Thus, to a hypothetical observer on Mercury, the Sun appears to move in a retrograde direction. Four Earth days after perihelion, the Sun's normal apparent motion resumes.[23] A similar effect would have occurred if Mercury had been in synchronous rotation: the alternating gain and loss of rotation over revolution would have caused a libration of 23.65° in longitude.[104]
	
	For the same reason, there are two points on Mercury's equator, 180 degrees apart in longitude, at either of which, around perihelion in alternate Mercurian years (once a Mercurian day), the Sun passes overhead, then reverses its apparent motion and passes overhead again, then reverses a second time and passes overhead a third time, taking a total of about 16 Earth-days for this entire process. In the other alternate Mercurian years, the same thing happens at the other of these two points. The amplitude of the retrograde motion is small, so the overall effect is that, for two or three weeks, the Sun is almost stationary overhead, and is at its most brilliant because Mercury is at perihelion, its closest to the Sun. This prolonged exposure to the Sun at its brightest makes these two points the hottest places on Mercury. Maximum temperature occurs when the Sun is at an angle of about 25 degrees past noon due to diurnal temperature lag, at 0.4 Mercury days and 0.8 Mercury years past sunrise.[105] Conversely, there are two other points on the equator, 90 degrees of longitude apart from the first ones, where the Sun passes overhead only when the planet is at aphelion in alternate years, when the apparent motion of the Sun in Mercury's sky is relatively rapid. These points, which are the ones on the equator where the apparent retrograde motion of the Sun happens when it is crossing the horizon as described in the preceding paragraph, receive much less solar heat than the first ones described above.[106]
	
	Mercury attains inferior conjunction (nearest approach to Earth) every 116 Earth days on average,[3] but this interval can range from 105 days to 129 days due to the planet's eccentric orbit. Mercury can come as near as 82,200,000 kilometres (0.549 astronomical units; 51.1 million miles) to Earth, and that is slowly declining: The next approach to within 82,100,000 km (51.0 million miles) is in 2679, and to within 82,000,000 km (51 million miles) in 4487, but it will not be closer to Earth than 80,000,000 km (50 million miles) until 28,622.[107] Its period of retrograde motion as seen from Earth can vary from 8 to 15 days on either side of inferior conjunction. This large range arises from the planet's high orbital eccentricity.[23] Essentially because Mercury is closest to the Sun, when taking an average over time, Mercury is the closest planet to the Earth,[108] and—in that measure—it is the closest planet to each of the other planets in the Solar System.[109][110][111][d]
	Longitude convention
	
	The longitude convention for Mercury puts the zero of longitude at one of the two hottest points on the surface, as described above. However, when this area was first visited, by Mariner 10, this zero meridian was in darkness, so it was impossible to select a feature on the surface to define the exact position of the meridian. Therefore, a small crater further west was chosen, called Hun Kal, which provides the exact reference point for measuring longitude.[112][113] The center of Hun Kal defines the 20° west meridian. A 1970 International Astronomical Union resolution suggests that longitudes be measured positively in the westerly direction on Mercury.[114] The two hottest places on the equator are therefore at longitudes 0° W and 180° W, and the coolest points on the equator are at longitudes 90° W and 270° W. However, the MESSENGER project uses an east-positive convention.[115]
	Spin-orbit resonance
	After one orbit, Mercury has rotated 1.5 times, so after two complete orbits the same hemisphere is again illuminated.
	
	For many years it was thought that Mercury was synchronously tidally locked with the Sun, rotating once for each orbit and always keeping the same face directed towards the Sun, in the same way that the same side of the Moon always faces Earth. Radar observations in 1965 proved that the planet has a 3:2 spin-orbit resonance, rotating three times for every two revolutions around the Sun. The eccentricity of Mercury's orbit makes this resonance stable—at perihelion, when the solar tide is strongest, the Sun is nearly still in Mercury's sky.[116]
	
	The rare 3:2 resonant tidal locking is stabilized by the variance of the tidal force along Mercury's eccentric orbit, acting on a permanent dipole component of Mercury's mass distribution.[117] In a circular orbit there is no such variance, so the only resonance stabilized in such an orbit is at 1:1 (e.g., Earth–Moon), when the tidal force, stretching a body along the "center-body" line, exerts a torque that aligns the body's axis of least inertia (the "longest" axis, and the axis of the aforementioned dipole) to point always at the center. However, with noticeable eccentricity, like that of Mercury's orbit, the tidal force has a maximum at perihelion and therefore stabilizes resonances, like 3:2, ensuring that the planet points its axis of least inertia roughly at the Sun when passing through perihelion.[117]
	
	The original reason astronomers thought it was synchronously locked was that, whenever Mercury was best placed for observation, it was always nearly at the same point in its 3:2 resonance, hence showing the same face. This is because, coincidentally, Mercury's rotation period is almost exactly half of its synodic period with respect to Earth. Due to Mercury's 3:2 spin-orbit resonance, a solar day lasts about 176 Earth days.[23] A sidereal day (the period of rotation) lasts about 58.7 Earth days.[23]
	
	Simulations indicate that the orbital eccentricity of Mercury varies chaotically from nearly zero (circular) to more than 0.45 over millions of years due to perturbations from the other planets.[23][118] This was thought to explain Mercury's 3:2 spin-orbit resonance (rather than the more usual 1:1), because this state is more likely to arise during a period of high eccentricity.[119] However, accurate modeling based on a realistic model of tidal response has demonstrated that Mercury was captured into the 3:2 spin-orbit state at a very early stage of its history, within 20 (more likely, 10) million years after its formation.[120]
	
	Numerical simulations show that a future secular orbital resonant perihelion interaction with Jupiter may cause the eccentricity of Mercury's orbit to increase to the point where there is a 1% chance that the planet will collide with Venus within the next five billion years.[121][122]
	Advance of perihelion
	Main article: Perihelion precession of Mercury
	
	In 1859, the French mathematician and astronomer Urbain Le Verrier reported that the slow precession of Mercury's orbit around the Sun could not be completely explained by Newtonian mechanics and perturbations by the known planets. He suggested, among possible explanations, that another planet (or perhaps instead a series of smaller 'corpuscules') might exist in an orbit even closer to the Sun than that of Mercury, to account for this perturbation.[123] (Other explanations considered included a slight oblateness of the Sun.) The success of the search for Neptune based on its perturbations of the orbit of Uranus led astronomers to place faith in this possible explanation, and the hypothetical planet was named Vulcan, but no such planet was ever found.[124]
	
	The perihelion precession of Mercury is 5,600 arcseconds (1.5556°) per century relative to Earth, or 574.10±0.65 arcseconds per century[125] relative to the inertial ICRF. Newtonian mechanics, taking into account all the effects from the other planets, predicts a precession of 5,557 arcseconds (1.5436°) per century.[125] In the early 20th century, Albert Einstein's general theory of relativity provided the explanation for the observed precession, by formalizing gravitation as being mediated by the curvature of spacetime. The effect is small: just 42.98 arcseconds per century for Mercury; it therefore requires a little over twelve million orbits for a full excess turn. Similar, but much smaller, effects exist for other Solar System bodies: 8.62 arcseconds per century for Venus, 3.84 for Earth, 1.35 for Mars, and 10.05 for 1566 Icarus.[126][127]
	Habitability
	See also: Mercury in fiction
	
	There may be scientific support, based on studies reported in March 2020, for considering that parts of the planet Mercury may have been habitable, and perhaps that life forms, albeit likely primitive microorganisms, may have existed on the planet.[128][129]
	Observation
	Image mosaic by Mariner 10, 1974
	
	Mercury's apparent magnitude is calculated to vary between −2.48 (brighter than Sirius) around superior conjunction and +7.25 (below the limit of naked-eye visibility) around inferior conjunction.[14] The mean apparent magnitude is 0.23 while the standard deviation of 1.78 is the largest of any planet. The mean apparent magnitude at superior conjunction is −1.89 while that at inferior conjunction is +5.93.[14] Observation of Mercury is complicated by its proximity to the Sun, as it is lost in the Sun's glare for much of the time. Mercury can be observed for only a brief period during either morning or evening twilight.[130]
	
	Mercury can, like several other planets and the brightest stars, be seen during a total solar eclipse.[131]
	
	Like the Moon and Venus, Mercury exhibits phases as seen from Earth. It is "new" at inferior conjunction and "full" at superior conjunction. The planet is rendered invisible from Earth on both of these occasions because of its being obscured by the Sun,[130] except its new phase during a transit.
	
	Mercury is technically brightest as seen from Earth when it is at a full phase. Although Mercury is farthest from Earth when it is full, the greater illuminated area that is visible and the opposition brightness surge more than compensates for the distance.[132] The opposite is true for Venus, which appears brightest when it is a crescent, because it is much closer to Earth than when gibbous.[132][133]
	False-color map showing the maximum temperatures of the north polar region
	
	Nonetheless, the brightest (full phase) appearance of Mercury is an essentially impossible time for practical observation, because of the extreme proximity of the Sun. Mercury is best observed at the first and last quarter, although they are phases of lesser brightness. The first and last quarter phases occur at greatest elongation east and west of the Sun, respectively. At both of these times Mercury's separation from the Sun ranges anywhere from 17.9° at perihelion to 27.8° at aphelion.[134][135] At greatest western elongation, Mercury rises at its earliest before sunrise, and at greatest eastern elongation, it sets at its latest after sunset.[136]
	False-color image of Carnegie Rupes, a tectonic landform—high terrain (red); low (blue).
	
	Mercury is more often and easily visible from the Southern Hemisphere than from the Northern. This is because Mercury's maximum western elongation occurs only during early autumn in the Southern Hemisphere, whereas its greatest eastern elongation happens only during late winter in the Southern Hemisphere.[136] In both of these cases, the angle at which the planet's orbit intersects the horizon is maximized, allowing it to rise several hours before sunrise in the former instance and not set until several hours after sundown in the latter from southern mid-latitudes, such as Argentina and South Africa.[136]
	
	An alternate method for viewing Mercury involves observing the planet during daylight hours when conditions are clear, ideally when it is at its greatest elongation. This allows the planet to be found easily, even when using telescopes with 8 cm (3.1 in) apertures. However, great care must be taken to obstruct the Sun from sight because of the extreme risk for eye damage.[137] This method bypasses the limitation of twilight observing when the ecliptic is located at a low elevation (e.g. on autumn evenings).
	
	Ground-based telescope observations of Mercury reveal only an illuminated partial disk with limited detail. The first of two spacecraft to visit the planet was Mariner 10, which mapped about 45% of its surface from 1974 to 1975. The second is the MESSENGER spacecraft, which after three Mercury flybys between 2008 and 2009, attained orbit around Mercury on March 17, 2011,[138] to study and map the rest of the planet.[139]
	
	The Hubble Space Telescope cannot observe Mercury at all, due to safety procedures that prevent its pointing too close to the Sun.[140]
	
	Because the shift of 0.15 revolutions in a year makes up a seven-year cycle (0.15 × 7 ≈ 1.0), in the seventh year Mercury follows almost exactly (earlier by 7 days) the sequence of phenomena it showed seven years before.[134]
	Observation history
	Ancient astronomers
	Mercury, from Liber astronomiae, 1550
	
	The earliest known recorded observations of Mercury are from the Mul.Apin tablets. These observations were most likely made by an Assyrian astronomer around the 14th century BC.[141] The cuneiform name used to designate Mercury on the Mul.Apin tablets is transcribed as Udu.Idim.Gu\u4.Ud ("the jumping planet").[e][142] Babylonian records of Mercury date back to the 1st millennium BC. The Babylonians called the planet Nabu after the messenger to the gods in their mythology.[143]
	
	The ancients knew Mercury by different names depending on whether it was an evening star or a morning star. By about 350 BC, the ancient Greeks had realized the two stars were one.[144] They knew the planet as Στίλβων Stilbōn, meaning "twinkling", and Ἑρμής Hermēs, for its fleeting motion,[145] a name that is retained in modern Greek (Ερμής Ermis).[146] The Romans named the planet after the swift-footed Roman messenger god, Mercury (Latin Mercurius), which they equated with the Greek Hermes, because it moves across the sky faster than any other planet.[144][147] The astronomical symbol for Mercury is a stylized version of Hermes' caduceus.[148]
	
	The Greco-Egyptian[149] astronomer Ptolemy wrote about the possibility of planetary transits across the face of the Sun in his work Planetary Hypotheses. He suggested that no transits had been observed either because planets such as Mercury were too small to see, or because the transits were too infrequent.[150]
	Ibn al-Shatir's model for the appearances of Mercury, showing the multiplication of epicycles using the Tusi couple, thus eliminating the Ptolemaic eccentrics and equant.
	
	In ancient China, Mercury was known as "the Hour Star" (Chen-xing 辰星). It was associated with the direction north and the phase of water in the Five Phases system of metaphysics.[151] Modern Chinese, Korean, Japanese and Vietnamese cultures refer to the planet literally as the "water star" (水星), based on the Five elements.[152][153][154] Hindu mythology used the name Budha for Mercury, and this god was thought to preside over Wednesday.[155] The god Odin (or Woden) of Germanic paganism was associated with the planet Mercury and Wednesday.[156] The Maya may have represented Mercury as an owl (or possibly four owls; two for the morning aspect and two for the evening) that served as a messenger to the underworld.[157]
	
	In medieval Islamic astronomy, the Andalusian astronomer Abū Ishāq Ibrāhīm al-Zarqālī in the 11th century described the deferent of Mercury's geocentric orbit as being oval, like an egg or a pignon, although this insight did not influence his astronomical theory or his astronomical calculations.[158][159] In the 12th century, Ibn Bajjah observed "two planets as black spots on the face of the Sun", which was later suggested as the transit of Mercury and/or Venus by the Maragha astronomer Qotb al-Din Shirazi in the 13th century.[160] (Note that most such medieval reports of transits were later taken as observations of sunspots.[161])
	
	In India, the Kerala school astronomer Nilakantha Somayaji in the 15th century developed a partially heliocentric planetary model in which Mercury orbits the Sun, which in turn orbits Earth, similar to the Tychonic system later proposed by Tycho Brahe in the late 16th century.[162]
	Ground-based telescopic research
	Transit of Mercury. Mercury is visible as a black dot below and to the left of center. The dark area above the center of the solar disk is a sunspot.
	Elongation is the angle between the Sun and the planet, with Earth as the reference point. Mercury appears close to the Sun.
	
	The first telescopic observations of Mercury were made by Galileo in the early 17th century. Although he observed phases when he looked at Venus, his telescope was not powerful enough to see the phases of Mercury. In 1631, Pierre Gassendi made the first telescopic observations of the transit of a planet across the Sun when he saw a transit of Mercury predicted by Johannes Kepler. In 1639, Giovanni Zupi used a telescope to discover that the planet had orbital phases similar to Venus and the Moon. The observation demonstrated conclusively that Mercury orbited around the Sun.[23]
	
	A rare event in astronomy is the passage of one planet in front of another (occultation), as seen from Earth. Mercury and Venus occult each other every few centuries, and the event of May 28, 1737 is the only one historically observed, having been seen by John Bevis at the Royal Greenwich Observatory.[163] The next occultation of Mercury by Venus will be on December 3, 2133.[164]
	
	The difficulties inherent in observing Mercury mean that it has been far less studied than the other planets. In 1800, Johann Schröter made observations of surface features, claiming to have observed 20-kilometre-high (12 mi) mountains. Friedrich Bessel used Schröter's drawings to erroneously estimate the rotation period as 24 hours and an axial tilt of 70°.[165] In the 1880s, Giovanni Schiaparelli mapped the planet more accurately, and suggested that Mercury's rotational period was 88 days, the same as its orbital period due to tidal locking.[166] This phenomenon is known as synchronous rotation. The effort to map the surface of Mercury was continued by Eugenios Antoniadi, who published a book in 1934 that included both maps and his own observations.[90] Many of the planet's surface features, particularly the albedo features, take their names from Antoniadi's map.[167]
	
	In June 1962, Soviet scientists at the Institute of Radio-engineering and Electronics of the USSR Academy of Sciences, led by Vladimir Kotelnikov, became the first to bounce a radar signal off Mercury and receive it, starting radar observations of the planet.[168][169][170] Three years later, radar observations by Americans Gordon H. Pettengill and Rolf B. Dyce, using the 300-meter Arecibo radio telescope in Puerto Rico, showed conclusively that the planet's rotational period was about 59 days.[171][172] The theory that Mercury's rotation was synchronous had become widely held, and it was a surprise to astronomers when these radio observations were announced. If Mercury were tidally locked, its dark face would be extremely cold, but measurements of radio emission revealed that it was much hotter than expected. Astronomers were reluctant to drop the synchronous rotation theory and proposed alternative mechanisms such as powerful heat-distributing winds to explain the observations.[173]
	Water ice (yellow) at Mercury's north polar region
	
	Italian astronomer Giuseppe Colombo noted that the rotation value was about two-thirds of Mercury's orbital period, and proposed that the planet's orbital and rotational periods were locked into a 3:2 rather than a 1:1 resonance.[174] Data from Mariner 10 subsequently confirmed this view.[175] This means that Schiaparelli's and Antoniadi's maps were not "wrong". Instead, the astronomers saw the same features during every second orbit and recorded them, but disregarded those seen in the meantime, when Mercury's other face was toward the Sun, because the orbital geometry meant that these observations were made under poor viewing conditions.[165]
	
	Ground-based optical observations did not shed much further light on Mercury, but radio astronomers using interferometry at microwave wavelengths, a technique that enables removal of the solar radiation, were able to discern physical and chemical characteristics of the subsurface layers to a depth of several meters.[176][177] Not until the first space probe flew past Mercury did many of its most fundamental morphological properties become known. Moreover, recent technological advances have led to improved ground-based observations. In 2000, high-resolution lucky imaging observations were conducted by the Mount Wilson Observatory 1.5 meter Hale telescope. They provided the first views that resolved surface features on the parts of Mercury that were not imaged in the Mariner 10 mission.[178] Most of the planet has been mapped by the Arecibo radar telescope, with 5 km (3.1 mi) resolution, including polar deposits in shadowed craters of what may be water ice.[179]
	Research with space probes
	Main article: Exploration of Mercury
	MESSENGER being prepared for launch
	Mercury transiting the Sun as viewed by the Mars rover Curiosity (June 3, 2014).[180]
	
	Reaching Mercury from Earth poses significant technical challenges, because it orbits so much closer to the Sun than Earth. A Mercury-bound spacecraft launched from Earth must travel over 91 million kilometres (57 million miles) into the Sun's gravitational potential well. Mercury has an orbital speed of 47.4 km/s (29.5 mi/s), whereas Earth's orbital speed is 29.8 km/s (18.5 mi/s).[100] Therefore, the spacecraft must make a large change in velocity (delta-v) to get to Mercury and then enter orbit,[181] as compared to the delta-v required for, say, Mars planetary missions.
	
	The potential energy liberated by moving down the Sun's potential well becomes kinetic energy, requiring a delta-v change to do anything other than pass by Mercury. Some portion of this delta-v budget can be provided from a gravity assist during one or more fly-bys of Venus.[182] To land safely or enter a stable orbit the spacecraft would rely entirely on rocket motors. Aerobraking is ruled out because Mercury has a negligible atmosphere. A trip to Mercury requires more rocket fuel than that required to escape the Solar System completely. As a result, only two space probes have visited it so far.[183] A proposed alternative approach would use a solar sail to attain a Mercury-synchronous orbit around the Sun.[184]
	Mariner 10
	Main article: Mariner 10
	Mariner 10, the first probe to visit Mercury
	
	The first spacecraft to visit Mercury was NASA's Mariner 10 (1974–1975).[144] The spacecraft used the gravity of Venus to adjust its orbital velocity so that it could approach Mercury, making it both the first spacecraft to use this gravitational "slingshot" effect and the first NASA mission to visit multiple planets.[185] Mariner 10 provided the first close-up images of Mercury's surface, which immediately showed its heavily cratered nature, and revealed many other types of geological features, such as the giant scarps that were later ascribed to the effect of the planet shrinking slightly as its iron core cools.[186] Unfortunately, the same face of the planet was lit at each of Mariner 10's close approaches. This made close observation of both sides of the planet impossible,[187] and resulted in the mapping of less than 45% of the planet's surface.[188]
	
	The spacecraft made three close approaches to Mercury, the closest of which took it to within 327 km (203 mi) of the surface.[189] At the first close approach, instruments detected a magnetic field, to the great surprise of planetary geologists—Mercury's rotation was expected to be much too slow to generate a significant dynamo effect. The second close approach was primarily used for imaging, but at the third approach, extensive magnetic data were obtained. The data revealed that the planet's magnetic field is much like Earth's, which deflects the solar wind around the planet. For many years after the Mariner 10 encounters, the origin of Mercury's magnetic field remained the subject of several competing theories.[190][191]
	
	On March 24, 1975, just eight days after its final close approach, Mariner 10 ran out of fuel. Because its orbit could no longer be accurately controlled, mission controllers instructed the probe to shut down.[192] Mariner 10 is thought to be still orbiting the Sun, passing close to Mercury every few months.[193]
	MESSENGER
	Main article: MESSENGER
	Estimated details of the impact of MESSENGER on April 30, 2015
	
	A second NASA mission to Mercury, named MESSENGER (MErcury Surface, Space ENvironment, GEochemistry, and Ranging), was launched on August 3, 2004. It made a fly-by of Earth in August 2005, and of Venus in October 2006 and June 2007 to place it onto the correct trajectory to reach an orbit around Mercury.[194] A first fly-by of Mercury occurred on January 14, 2008, a second on October 6, 2008,[195] and a third on September 29, 2009.[196] Most of the hemisphere not imaged by Mariner 10 was mapped during these fly-bys. The probe successfully entered an elliptical orbit around the planet on March 18, 2011. The first orbital image of Mercury was obtained on March 29, 2011. The probe finished a one-year mapping mission,[195] and then entered a one-year extended mission into 2013. In addition to continued observations and mapping of Mercury, MESSENGER observed the 2012 solar maximum.[197]
	
	The mission was designed to clear up six key issues: Mercury's high density, its geological history, the nature of its magnetic field, the structure of its core, whether it has ice at its poles, and where its tenuous atmosphere comes from. To this end, the probe carried imaging devices that gathered much-higher-resolution images of much more of Mercury than Mariner 10, assorted spectrometers to determine abundances of elements in the crust, and magnetometers and devices to measure velocities of charged particles. Measurements of changes in the probe's orbital velocity were expected to be used to infer details of the planet's interior structure.[198] MESSENGER's final maneuver was on April 24, 2015, and it crashed into Mercury's surface on April 30, 2015.[199][200][201] The spacecraft's impact with Mercury occurred near 3:26 PM EDT on April 30, 2015, leaving a crater estimated to be 16 m (52 ft) in diameter.[202]
	BepiColombo
	Main article: BepiColombo
	
	The European Space Agency and the Japanese Space Agency developed and launched a joint mission called BepiColombo, which will orbit Mercury with two probes: one to map the planet and the other to study its magnetosphere.[203] Launched on October 20, 2018, BepiColombo is expected to reach Mercury in 2025.[204] It will release a magnetometer probe into an elliptical orbit, then chemical rockets will fire to deposit the mapper probe into a circular orbit. Both probes will operate for one terrestrial year.[203] The mapper probe carries an array of spectrometers similar to those on MESSENGER, and will study the planet at many different wavelengths including infrared, ultraviolet, X-ray and gamma ray.[205]
	Comparison
	Size comparison with other Solar System objects
	Mercury, Earth
	Mercury, Venus, Earth, Mars
	Back row: Mars, Mercury
	Front: Moon, Pluto, Haumea
	See also
	
		Outline of Mercury (planet)
		Budha, Hinduism's name for the planet and the god Mercury
		Colonization of Mercury
		Mercury in astrology
		Mercury in fiction
	
	Notes
	
	In astronomy, the words "rotation" and "revolution" have different meanings. "Rotation" is the turning of a body about an axis that passes through the body, as in "Earth rotates once a day." "Revolution" is motion around a centre that is external to the body, usually in orbit, as in "Earth takes a year for each revolution around the Sun." The verbs "rotate" and "revolve" mean doing rotation and revolution, respectively.
	Pluto was considered a planet from its discovery in 1930 to 2006, but after that it has been reclassified as a dwarf planet. Pluto's orbital eccentricity is greater than Mercury's. Pluto is also smaller than Mercury, but was thought to be larger until 1976.
	The Sun's total angular displacement during its apparent retrograde motion as seen from the surface of Mercury is ~1.23°, while the Sun's angular diameter when the apparent retrograde motion begins and ends is ~1.71°, increasing to ~1.73° at perihelion (midway through the retrograde motion).
	It is important to be clear about the meaning of 'closeness'. In the astronomical literature, the term 'closest planets' often means 'the two planets that approach each other most closely'. In other words, the orbits of the two planets approach each other most closely. However, this does not mean that the two planets are closest over time. For example, essentially because Mercury is closer to the Sun than Venus, Mercury spends more time in proximity to Earth; it could, therefore, be said that Mercury is the planet that is 'closest to Earth when averaged over time'. However, using this time-average definition of 'closeness'—as noted above—it turns out that Mercury is the closest planet to all other planets in the solar system. For that reason, arguably, the proximity-definition is not particularly helpful. An episode of the BBC Radio 4 programme 'More or Less' explains the different notions of proximity well.[108]
	
		Some sources precede the cuneiform transcription with "MUL". "MUL" is a cuneiform sign that was used in the Sumerian language to designate a star or planet, but it is not considered part of the actual name. The "4" is a reference number in the Sumero–Akkadian transliteration system to designate which of several syllables a certain cuneiform sign is most likely designating.
	
	References
	
	"Mercurian". Lexico UK Dictionary. Oxford University Press.
	"Mercurial". Lexico UK Dictionary. Oxford University Press.
	Williams, David R. (November 25, 2020). "Mercury Fact Sheet". NASA. Retrieved April 19, 2021.
	Souami, D.; Souchay, J. (July 2012). "The solar system's invariable plane". Astronomy & Astrophysics. 543: 11. Bibcode:2012A&A...543A.133S. doi:10.1051/0004-6361/201219011. A133.
	Yeomans, Donald K. (April 7, 2008). "HORIZONS Web-Interface for Mercury Major Body". JPL Horizons On-Line Ephemeris System. Retrieved April 7, 2008. – Select "Ephemeris Type: Orbital Elements", "Time Span: 2000-01-01 12:00 to 2000-01-02". ("Target Body: Mercury" and "Center: Sun" should be defaulted to.) Results are instantaneous osculating values at the precise J2000 epoch.
	Davis, Phillips; Barnett, Amanda (February 15, 2021). "Mercury". Solar System Exploration. NASA Jet Propulsion Laboratory. Retrieved April 21, 2021.
	Seidelmann, P. Kenneth; Archinal, Brent A.; A'Hearn, Michael F.; et al. (2007). "Report of the IAU/IAG Working Group on cartographic coordinates and rotational elements: 2006". Celestial Mechanics and Dynamical Astronomy. 98 (3): 155–180. Bibcode:2007CeMDA..98..155S. doi:10.1007/s10569-007-9072-y. S2CID 122772353.
	Mazarico, Erwan; Genova, Antonio; Goossens, Sander; Lemoine, Frank G.; Neumann, Gregory A.; Zuber, Maria T.; Smith, David E.; Solomon, Sean C. (2014). "The gravity field, orientation, and ephemeris of Mercury from MESSENGER observations after three years in orbit" (PDF). Journal of Geophysical Research: Planets. 119 (12): 2417–2436. Bibcode:2014JGRE..119.2417M. doi:10.1002/2014JE004675. hdl:1721.1/97927. ISSN 2169-9097.
	Margot, Jean-Luc; Peale, Stanton J.; Solomon, Sean C.; Hauck, Steven A.; Ghigo, Frank D.; Jurgens, Raymond F.; Yseboodt, Marie; Giorgini, Jon D.; Padovan, Sebastiano; Campbell, Donald B. (2012). "Mercury's moment of inertia from spin and gravity data". Journal of Geophysical Research: Planets. 117 (E12): n/a. Bibcode:2012JGRE..117.0L09M. CiteSeerX 10.1.1.676.5383. doi:10.1029/2012JE004161. ISSN 0148-0227.
	"ESO". ESO. Retrieved June 3, 2021.
	Mallama, Anthony (2017). "The spherical bolometric albedo for planet Mercury". arXiv:1703.02670 [astro-ph.EP].
	Mallama, Anthony; Wang, Dennis; Howard, Russell A. (2002). "Photometry of Mercury from SOHO/LASCO and Earth". Icarus. 155 (2): 253–264. Bibcode:2002Icar..155..253M. doi:10.1006/icar.2001.6723.
	Vasavada, Ashwin R.; Paige, David A.; Wood, Stephen E. (February 19, 1999). "Near-Surface Temperatures on Mercury and the Moon and the Stability of Polar Ice Deposits" (PDF). Icarus. 141 (2): 179–193. Bibcode:1999Icar..141..179V. doi:10.1006/icar.1999.6175. Figure 3 with the "TWO model"; Figure 5 for pole.
	Mallama, Anthony; Hilton, James L. (October 2018). "Computing apparent planetary magnitudes for The Astronomical Almanac". Astronomy and Computing. 25: 10–24. arXiv:1808.01973. Bibcode:2018A&C....25...10M. doi:10.1016/j.ascom.2018.08.002. S2CID 69912809.
	Milillo, A.; Wurz, P.; Orsini, S.; Delcourt, D.; Kallio, E.; Killen, R. M.; Lammer, H.; Massetti, S.; Mura, A.; Barabash, S.; Cremonese, G.; Daglis, I. A.; Angelis, E.; Lellis, A. M.; Livi, S.; Mangano, V.; Torkar, K. (April 2005). "Surface-Exosphere-Magnetosphere System Of Mercury". Space Science Reviews. 117 (3–4): 397–443. Bibcode:2005SSRv..117..397M. doi:10.1007/s11214-005-3593-z. S2CID 122285073.
	Berezhnoy, Alexey A. (January 2018). "Chemistry of impact events on Mercury". Icarus. 300: 210–222. Bibcode:2018Icar..300..210B. doi:10.1016/j.icarus.2017.08.034.
	Elkins-Tanton, Linda T. (2006). Uranus, Neptune, Pluto, and the Outer Solar System. Infobase Publishing. p. 51. ISBN 978-1-4381-0729-5. Extract of page 51
	"Animated clip of orbit and rotation of Mercury". Sciencenetlinks.com.
	Prockter, Louise (2005). Ice in the Solar System (PDF). 26. Johns Hopkins APL Technical Digest. Retrieved July 27, 2009.
	"NASA Completes MESSENGER Mission with Expected Impact on Mercury's Surface". Archived from the original on May 3, 2015. Retrieved April 30, 2015.
	"From Mercury orbit, MESSENGER watches a lunar eclipse". Planetary Society. October 10, 2014. Retrieved January 23, 2015.
	"Innovative use of pressurant extends MESSENGER's Mercury mission". Astronomy.com. December 29, 2014. Retrieved January 22, 2015.
	Strom, Robert G.; Sprague, Ann L. (2003). Exploring Mercury: the iron planet. Springer. ISBN 978-1-85233-731-5.
	Talbert, Tricia, ed. (March 21, 2012). "MESSENGER Provides New Look at Mercury's Surprising Core and Landscape Curiosities". NASA.
	"Scientists find evidence Mercury has a solid inner core". AGU Newsroom. Retrieved April 17, 2019.
	"Mercury". US Geological Survey. May 8, 2003. Archived from the original on September 29, 2006. Retrieved November 26, 2006.
	Lyttleton, Raymond A. (1969). "On the Internal Structures of Mercury and Venus". Astrophysics and Space Science. 5 (1): 18–35. Bibcode:1969Ap&SS...5...18L. doi:10.1007/BF00653933. S2CID 122572625.
	Gold, Lauren (May 3, 2007). "Mercury has molten core, Cornell researcher shows". Chronicle Online. Cornell University. Retrieved May 12, 2008.
	Finley, Dave (May 3, 2007). "Mercury's Core Molten, Radar Study Shows". National Radio Astronomy Observatory. Retrieved May 12, 2008.
	Spohn, Tilman; Sohl, Frank; Wieczerkowski, Karin; Conzelmann, Vera (2001). "The interior structure of Mercury: what we know, what we expect from BepiColombo". Planetary and Space Science. 49 (14–15): 1561–1570. Bibcode:2001P&SS...49.1561S. doi:10.1016/S0032-0633(01)00093-9.
	Gallant, Roy A. (1986). The National Geographic Picture Atlas of Our Universe (2nd ed.). National Geographic Society. ISBN 9780870446443.
	Padovan, Sebastiano; Wieczorek, Mark A.; Margot, Jean-Luc; Tosi, Nicola; Solomon, Sean C. (2015). "Thickness of the crust of Mercury from geoid-to-topography ratios". Geophysical Research Letters. 42 (4): 1029. Bibcode:2015GeoRL..42.1029P. doi:10.1002/2014GL062487.
	Sori, Michael M. (May 2018). "A thin, dense crust for Mercury". Earth and Planetary Science Letters. 489: 92–99. Bibcode:2018E&PSL.489...92S. doi:10.1016/j.epsl.2018.02.033.
	Schenk, Paul M.; Melosh, H. Jay (March 1994). "Lobate Thrust Scarps and the Thickness of Mercury's Lithosphere". Abstracts of the 25th Lunar and Planetary Science Conference. 1994: 1994LPI....25.1203S. Bibcode:1994LPI....25.1203S.
	Watters, T. R.; Nimmo, F.; Robinson, M. S. (2004). Chronology of Lobate Scarp Thrust Faults and the Mechanical Structure of Mercury's Lithosphere. Lunar and Planetary Science Conference. p. 1886. Bibcode:2004LPI....35.1886W.
	Watters, Thomas R.; Robinson, Mark S.; Cook, Anthony C. (November 1998). "Topography of lobate scarps on Mercury; new constraints on the planet's contraction". Geology. 26 (11): 991–994. Bibcode:1998Geo....26..991W. doi:10.1130/0091-7613(1998)026<0991:TOLSOM>2.3.CO;2.
	Benz, W.; Slattery, W. L.; Cameron, Alastair G. W. (1988). "Collisional stripping of Mercury's mantle". Icarus. 74 (3): 516–528. Bibcode:1988Icar...74..516B. doi:10.1016/0019-1035(88)90118-2.
	Cameron, Alastair G. W. (1985). "The partial volatilization of Mercury". Icarus. 64 (2): 285–294. Bibcode:1985Icar...64..285C. doi:10.1016/0019-1035(85)90091-0.
	Weidenschilling, Stuart J. (1987). "Iron/silicate fractionation and the origin of Mercury". Icarus. 35 (1): 99–111. Bibcode:1978Icar...35...99W. doi:10.1016/0019-1035(78)90064-7.
	Sappenfield, Mark (September 29, 2011). "Messenger's message from Mercury: Time to rewrite the textbooks". The Christian Science Monitor. Retrieved August 21, 2017.
	"BepiColombo". Science & Technology. European Space Agency. Retrieved April 7, 2008.
	Cartwright, Jon (September 30, 2011). "Messenger sheds light on Mercury's formation". Chemistry World. Retrieved August 21, 2017.
	Morris, Jefferson (November 10, 2008). "Laser Altimetry". Aviation Week & Space Technology. 169 (18): 18. "Mercury's crust is more analogous to a marbled cake than a layered cake."
	Hughes, E. T.; Vaughan, W. M. (March 2012). Albedo Features of Mercury. 43rd Lunar and Planetary Science Conference, held March 19–23, 2012 at The Woodlands, Texas. 1659. Bibcode:2012LPI....43.2151H. 2151.
	Blue, Jennifer (April 11, 2008). "Gazetteer of Planetary Nomenclature". US Geological Survey. Retrieved April 11, 2008.
	Dunne, James A.; Burgess, Eric (1978). "Chapter Seven". The Voyage of Mariner 10 – Mission to Venus and Mercury. NASA History Office. Retrieved May 28, 2008.
	Nittler, Larry R.; Weider, Shoshana Z. (2019). "The Surface Composition of Mercury". Elements. 15 (1): 33–38. doi:10.2138/gselements.15.1.33.
	"Categories for Naming Features on Planets and Satellites". US Geological Survey. Retrieved August 20, 2011.
	Strom, Robert G. (1979). "Mercury: a post-Mariner assessment". Space Science Reviews. 24 (1): 3–70. Bibcode:1979SSRv...24....3S. doi:10.1007/BF00221842. S2CID 122563809.
	Broadfoot, A. Lyle; Kumar, Shailendra; Belton, Michael J. S.; McElroy, Michael B. (July 12, 1974). "Mercury's Atmosphere from Mariner 10: Preliminary Results". Science. 185 (4146): 166–169. Bibcode:1974Sci...185..166B. doi:10.1126/science.185.4146.166. PMID 17810510. S2CID 7790470.
	Geology of the solar system. IMAP 2596. U.S. Geological Survey. 1997. doi:10.3133/i2596.
	Head, James W.; Solomon, Sean C. (1981). "Tectonic Evolution of the Terrestrial Planets" (PDF). Science. 213 (4503): 62–76. Bibcode:1981Sci...213...62H. CiteSeerX 10.1.1.715.4402. doi:10.1126/science.213.4503.62. PMID 17741171.
	"Scientists see Mercury in a new light". Science Daily. February 28, 2008. Retrieved April 7, 2008.
	"The Giant Spider of Mercury". The Planetary Society. Retrieved June 9, 2017.
	Spudis, Paul D. (2001). "The Geological History of Mercury". Workshop on Mercury: Space Environment, Surface, and Interior, Chicago (1097): 100. Bibcode:2001mses.conf..100S.
	Ritzel, Rebecca (December 20, 2012). "Ballet isn't rocket science, but the two aren't mutually exclusive, either". Washington Post. Washington, D.C., United States. Retrieved December 22, 2012.
	Shiga, David (January 30, 2008). "Bizarre spider scar found on Mercury's surface". NewScientist.com news service.
	Schultz, Peter H.; Gault, Donald E. (1975). "Seismic effects from major basin formations on the moon and Mercury". Earth, Moon, and Planets. 12 (2): 159–175. Bibcode:1975Moon...12..159S. doi:10.1007/BF00577875. S2CID 121225801.
	Wieczorek, Mark A.; Zuber, Maria T. (2001). "A Serenitatis origin for the Imbrian grooves and South Pole-Aitken thorium anomaly". Journal of Geophysical Research. 106 (E11): 27853–27864. Bibcode:2001JGR...10627853W. doi:10.1029/2000JE001384. Retrieved May 12, 2008.
	Fassett, Caleb I.; Head, James W.; Baker, David M. H.; Zuber, Maria T.; Smith, David E.; Neumann, Gregory A.; Solomon, Sean C.; Klimczak, Christian; Strom, Robert G.; Chapman, Clark R.; Prockter, Louise M.; Phillips, Roger J.; Oberst, Jürgen; Preusker, Frank (October 2012). "Large impact basins on Mercury: Global distribution, characteristics, and modification history from MESSENGER orbital data". Journal of Geophysical Research. 117. 15 pp. Bibcode:2012JGRE..117.0L08F. doi:10.1029/2012JE004154. E00L08.
	Denevi, Brett W.; Robinson, Mark S. (2008). "Albedo of Immature Mercurian Crustal Materials: Evidence for the Presence of Ferrous Iron". Lunar and Planetary Science. 39 (1391): 1750. Bibcode:2008LPI....39.1750D.
	Wagner, Roland J.; Wolf, Ursula; Ivanov, Boris A.; Neukum, Gerhard (October 4–5, 2001). Application of an Updated Impact Cratering Chronology Model to Mercury' s Time-Stratigraphic System. Workshop on Mercury: Space Environment, Surface, and Interior. Proceedings of a workshop held at The Field Museum. Chicago, IL: Lunar and Planetary Science Institute. p. 106. Bibcode:2001mses.conf..106W.
	Choi, Charles Q. (September 26, 2016). "Mercuryquakes May Currently Shake Up the Tiny Planet". Space.com. Retrieved September 28, 2016.
	Dzurisin, Daniel (October 10, 1978). "The tectonic and volcanic history of Mercury as inferred from studies of scarps, ridges, troughs, and other lineaments". Journal of Geophysical Research. 83 (B10): 4883–4906. Bibcode:1978JGR....83.4883D. doi:10.1029/JB083iB10p04883.
	Watters, Thomas R.; Daud, Katie; Banks, Maria E.; Selvans, Michelle M.; Chapman, Clark R.; Ernst, Carolyn M. (September 26, 2016). "Recent tectonic activity on Mercury revealed by small thrust fault scarps". Nature Geoscience. 9 (10): 743–747. Bibcode:2016NatGe...9..743W. doi:10.1038/ngeo2814.
	Giacomini, L.; Massironi, M.; Galluzzi, V.; Ferrari, S.; Palumbo, P. (May 2020). "Dating long thrust systems on Mercury: New clues on the thermal evolution of the planet". Geoscience Frontiers. 11 (3): 855–870. doi:10.1016/j.gsf.2019.09.005.
	Schleicher, Lisa S.; Watters, Thomas R.; Martin, Aaron J.; Banks, Maria E. (October 2019). "Wrinkle ridges on Mercury and the Moon within and outside of mascons". Icarus. 331: 226–237. Bibcode:2019Icar..331..226S. doi:10.1016/j.icarus.2019.04.013.
	Kerber, Laura; Head, James W.; Solomon, Sean C.; Murchie, Scott L.; Blewett, David T. (August 15, 2009). "Explosive volcanic eruptions on Mercury: Eruption conditions, magma volatile content, and implications for interior volatile abundances". Earth and Planetary Science Letters. 285 (3–4): 263–271. Bibcode:2009E&PSL.285..263K. doi:10.1016/j.epsl.2009.04.037.
	Head, James W.; Chapman, Clark R.; Strom, Robert G.; Fassett, Caleb I.; Denevi, Brett W. (September 30, 2011). "Flood Volcanism in the Northern High Latitudes of Mercury Revealed by MESSENGER" (PDF). Science. 333 (6051): 1853–1856. Bibcode:2011Sci...333.1853H. doi:10.1126/science.1211997. PMID 21960625. S2CID 7651992.
	Thomas, Rebecca J.; Rothery, David A.; Conway, Susan J.; Anand, Mahesh (September 16, 2014). "Long-lived explosive volcanism on Mercury". Geophysical Research Letters. 41 (17): 6084–6092. Bibcode:2014GeoRL..41.6084T. doi:10.1002/2014GL061224.
	Groudge, Timothy A.; Head, James W. (March 2014). "Global inventory and characterization of pyroclastic deposits on Mercury: New insights into pyroclastic activity from MESSENGER orbital data" (PDF). Journal of Geophysical Research. 119 (3): 635–658. Bibcode:2014JGRE..119..635G. doi:10.1002/2013JE004480.
	Rothery, David A.; Thomas, Rebeca J.; Kerber, Laura (January 1, 2014). "Prolonged eruptive history of a compound volcano on Mercury: Volcanic and tectonic implications" (PDF). Earth and Planetary Science Letters. 385: 59–67. Bibcode:2014E&PSL.385...59R. doi:10.1016/j.epsl.2013.10.023.
	Chang, Kenneth (November 29, 2012). "On Closest Planet to the Sun, NASA Finds Lots of Ice". The New York Times. p. A3. Archived from the original on November 29, 2012. "Sean C. Solomon, the principal investigator for MESSENGER, said there was enough ice there to encase Washington, D.C., in a frozen block two and a half miles deep."
	Lewis, John S. (2004). Physics and Chemistry of the Solar System (2nd ed.). Academic Press. p. 463. ISBN 978-0-12-446744-6.
	Murdock, Thomas L.; Ney, Edward P. (1970). "Mercury: The Dark-Side Temperature". Science. 170 (3957): 535–537. Bibcode:1970Sci...170..535M. doi:10.1126/science.170.3957.535. PMID 17799708. S2CID 38824994.
	Lewis, John S. (2004). Physics and Chemistry of the Solar System. Academic Press. ISBN 978-0-12-446744-6. Retrieved June 3, 2008.
	Ingersoll, Andrew P.; Svitek, Tomas; Murray, Bruce C. (1992). "Stability of polar frosts in spherical bowl-shaped craters on the Moon, Mercury, and Mars". Icarus. 100 (1): 40–47. Bibcode:1992Icar..100...40I. doi:10.1016/0019-1035(92)90016-Z.
	Slade, Martin A.; Butler, Bryan J.; Muhleman, Duane O. (1992). "Mercury radar imaging – Evidence for polar ice". Science. 258 (5082): 635–640. Bibcode:1992Sci...258..635S. doi:10.1126/science.258.5082.635. PMID 17748898. S2CID 34009087.
	Williams, David R. (June 2, 2005). "Ice on Mercury". NASA Goddard Space Flight Center. Retrieved May 23, 2008.
	Rawlins, Katherine; Moses, Julianne I.; Zahnle, Kevin J. (1995). "Exogenic Sources of Water for Mercury's Polar Ice". Bulletin of the American Astronomical Society. 27: 1117. Bibcode:1995DPS....27.2112R.
	Harmon, John K.; Perillat, Phil J.; Slade, Martin A. (2001). "High-Resolution Radar Imaging of Mercury's North Pole". Icarus. 149 (1): 1–15. Bibcode:2001Icar..149....1H. doi:10.1006/icar.2000.6544.
	Domingue DL, Koehn PL, et al. (2009). "Mercury's Atmosphere: A Surface-Bounded Exosphere". Space Science Reviews. 131 (1–4): 161–186. Bibcode:2007SSRv..131..161D. doi:10.1007/s11214-007-9260-9. S2CID 121301247.
	Hunten, Donald M.; Shemansky, Donald Eugene; Morgan, Thomas Hunt (1988). "The Mercury atmosphere". In Vilas, Faith; Chapman, Clark R.; Shapley Matthews, Mildred (eds.). Mercury. University of Arizona Press. ISBN 978-0-8165-1085-6.
	Lakdawalla, Emily (July 3, 2008). "MESSENGER Scientists 'Astonished' to Find Water in Mercury's Thin Atmosphere". The Planetary Society. Retrieved May 18, 2009.
	Zurbuchen TH, Raines JM, et al. (2008). "MESSENGER Observations of the Composition of Mercury's Ionized Exosphere and Plasma Environment". Science. 321 (5885): 90–92. Bibcode:2008Sci...321...90Z. doi:10.1126/science.1159314. PMID 18599777. S2CID 206513512.
	"Instrument Shows What Planet Mercury Is Made Of". University of Michigan. June 30, 2008. Retrieved May 18, 2009.
	Killen, Rosemary; Cremonese, Gabrielle; et al. (2007). "Processes that Promote and Deplete the Exosphere of Mercury". Space Science Reviews. 132 (2–4): 433–509. Bibcode:2007SSRv..132..433K. doi:10.1007/s11214-007-9232-0. S2CID 121944553.
	Killen, Rosemary M.; Hahn, Joseph M. (December 10, 2014). "Impact Vaporization as a Possible Source of Mercury's Calcium Exosphere". Icarus. 250: 230–237. Bibcode:2015Icar..250..230K. doi:10.1016/j.icarus.2014.11.035. hdl:2060/20150010116.
	McClintock, William E.; Vervack, Ronald J.; et al. (2009). "MESSENGER Observations of Mercury's Exosphere: Detection of Magnesium and Distribution of Constituents". Science. 324 (5927): 610–613. Bibcode:2009Sci...324..610M. doi:10.1126/science.1172525 (inactive May 31, 2021). PMID 19407195.
	Beatty, J. Kelly; Petersen, Carolyn Collins; Chaikin, Andrew (1999). The New Solar System. Cambridge University Press. ISBN 978-0-521-64587-4.
	Seeds, Michael A. (2004). Astronomy: The Solar System and Beyond (4th ed.). Brooks Cole. ISBN 978-0-534-42111-3.
	Williams, David R. (January 6, 2005). "Planetary Fact Sheets". NASA National Space Science Data Center. Retrieved August 10, 2006.
	"Mercury's Internal Magnetic Field". NASA. January 30, 2008. Retrieved April 21, 2021.
	Gold, Lauren (May 3, 2007). "Mercury has molten core, Cornell researcher shows". Cornell University. Retrieved April 7, 2008.
	Christensen, Ulrich R. (2006). "A deep dynamo generating Mercury's magnetic field". Nature. 444 (7122): 1056–1058. Bibcode:2006Natur.444.1056C. doi:10.1038/nature05342. PMID 17183319. S2CID 4342216.
	Padovan, Sebastiano; Margot, Jean-Luc; Hauck, Steven A.; Moore, William B.; Solomon, Sean C. (April 2014). "The tides of Mercury and possible implications for its interior structure". Journal of Geophysical Research: Planets. 119 (4): 850–866. Bibcode:2014JGRE..119..850P. doi:10.1002/2013JE004459.
	Steigerwald, Bill (June 2, 2009). "Magnetic Tornadoes Could Liberate Mercury's Tenuous Atmosphere". NASA Goddard Space Flight Center. Retrieved July 18, 2009.
	Van Hoolst, Tim; Jacobs, Carla (2003). "Mercury's tides and interior structure". Journal of Geophysical Research. 108 (E11): 7. Bibcode:2003JGRE..108.5121V. doi:10.1029/2003JE002126.
	"Space Topics: Compare the Planets: Mercury, Venus, Earth, The Moon, and Mars". Planetary Society. Archived from the original on July 28, 2011. Retrieved April 12, 2007.
	Williams, David R. (October 21, 2019). "Planetary Fact Sheet - Metric". NASA. Retrieved April 20, 2021.
	Espenak, Fred (April 21, 2005). "Transits of Mercury". NASA/Goddard Space Flight Center. Retrieved May 20, 2008.
	Biswas, Sukumar (2000). Cosmic Perspectives in Space Physics. Astrophysics and Space Science Library. Springer. p. 176. ISBN 978-0-7923-5813-8.
	Margot, J. L.; Peale, S. J.; Jurgens, R. F.; Slade, M. A.; et al. (2007). "Large Longitude Libration of Mercury Reveals a Molten Core". Science. 316 (5825): 710–714. Bibcode:2007Sci...316..710M. doi:10.1126/science.1140514. PMID 17478713. S2CID 8863681.
	Popular Astronomy: A Review of Astronomy and Allied Sciences. Goodsell Observatory of Carleton College. 1896. "although in the case of Venus the libration in longitude due to the eccentricity of the orbit amounts to only 47' on either side of the mean position, in the case of Mercury it amounts to 23° 39'"
	Seligman, C. "The Rotation of Mercury". cseligman.com. NASA Flash animation. Retrieved July 31, 2019.
	van Hemerlrijck, E. (August 1983). "On the Variations in the Insolation at Mercury Resulting from Oscillations of the Orbital Eccentricity". The Moon and the Planets. 29 (1): 83–93. Bibcode:1983M&P....29...83V. doi:10.1007/BF00928377. S2CID 122761699.
	Mercury Closest Approaches to Earth generated with:
	1. Solex 10  Archived April 29, 2009, at WebCite (Text Output file Archived March 9, 2012, at the Wayback Machine)
	2. Gravity Simulator charts Archived September 12, 2014, at the Wayback Machine
	3. JPL Horizons 1950–2200  Archived November 6, 2015, at the Wayback Machine
	(3 sources are provided to address original research concerns and to support general long-term trends)
	Harford, Tim (January 11, 2019). "BBC Radio 4 – More or Less, Sugar, Outdoors Play and Planets". BBC. "Oliver Hawkins, more or less alumnus and statistical legend, wrote some code for us, which calculated which planet was closest to the Earth on each day for the past 50 years, and then sent the results to David A. Rothery, professor of planetary geosciences at the Open University."
	Stockman, Tom; Monroe, Gabriel; Cordner, Samuel (March 12, 2019). "Venus is not Earth's closest neighbor". Physics Today. doi:10.1063/PT.6.3.20190312a.
	Stockman, Tom (March 7, 2019). Mercury is the closest planet to all seven other planets (video). YouTube. Retrieved May 29, 2019.
	🌍 Which Planet is the Closest?, retrieved July 22, 2021
	Davies, M. E. (June 10, 1975). "Surface Coordinates and Cartography of Mercury". Journal of Geophysical Research. 80 (B17): 2417–2430. Bibcode:1975JGR....80.2417D. doi:10.1029/JB080i017p02417.
	Davies, M. E.; Dwornik, S. E.; Gault, D. E.; Strom, R. G. (1978). NASA Atlas of Mercury. NASA Scientific and Technical Information Office.
	"USGS Astrogeology: Rotation and pole position for the Sun and planets (IAU WGCCRE)". Archived from the original on October 24, 2011. Retrieved October 22, 2009.
	Archinal, Brent A.; A'Hearn, Michael F.; Bowell, Edward L.; Conrad, Albert R.; et al. (2010). "Report of the IAU Working Group on Cartographic Coordinates and Rotational Elements: 2009". Celestial Mechanics and Dynamical Astronomy. 109 (2): 101–135. Bibcode:2011CeMDA.109..101A. doi:10.1007/s10569-010-9320-4. ISSN 0923-2958. S2CID 189842666.
	Liu, Han-Shou; O'Keefe, John A. (1965). "Theory of Rotation for the Planet Mercury". Science. 150 (3704): 1717. Bibcode:1965Sci...150.1717L. doi:10.1126/science.150.3704.1717. PMID 17768871. S2CID 45608770.
	Colombo, Giuseppe; Shapiro, Irwin I. (1966). "The rotation of the planet Mercury". Astrophysical Journal. 145: 296. Bibcode:1966ApJ...145..296C. doi:10.1086/148762.
	Correia, Alexandre C. M.; Laskar, Jacques (2009). "Mercury's capture into the 3/2 spin-orbit resonance including the effect of core–mantle friction". Icarus. 201 (1): 1–11. arXiv:0901.1843. Bibcode:2009Icar..201....1C. doi:10.1016/j.icarus.2008.12.034. S2CID 14778204.
	Correia, Alexandre C. M.; Laskar, Jacques (2004). "Mercury's capture into the 3/2 spin-orbit resonance as a result of its chaotic dynamics". Nature. 429 (6994): 848–850. Bibcode:2004Natur.429..848C. doi:10.1038/nature02609. PMID 15215857. S2CID 9289925.
	Noyelles, B.; Frouard, J.; Makarov, V. V. & Efroimsky, M. (2014). "Spin-orbit evolution of Mercury revisited". Icarus. 241 (2014): 26–44. arXiv:1307.0136. Bibcode:2014Icar..241...26N. doi:10.1016/j.icarus.2014.05.045. S2CID 53690707.
	Laskar, Jacques (March 18, 2008). "Chaotic diffusion in the Solar System". Icarus. 196 (1): 1–15. arXiv:0802.3371. Bibcode:2008Icar..196....1L. doi:10.1016/j.icarus.2008.02.017. S2CID 11586168.
	Laskar, Jacques; Gastineau, Mickaël (June 11, 2009). "Existence of collisional trajectories of Mercury, Mars and Venus with the Earth". Nature. 459 (7248): 817–819. Bibcode:2009Natur.459..817L. doi:10.1038/nature08096. PMID 19516336. S2CID 4416436.
	Le Verrier, Urbain (1859). "Lettre de M. Le Verrier à M. Faye sur la théorie de Mercure et sur le mouvement du périhélie de cette planète". Comptes rendus hebdomadaires des séances de l'Académie des sciences (in French). Paris. 49: 379–383. (At p. 383 in the same volume Le Verrier's report is followed by another, from Faye, enthusiastically recommending to astronomers to search for a previously undetected intra-mercurial object.)
	Baum, Richard; Sheehan, William (1997). In Search of Planet Vulcan, The Ghost in Newton's Clockwork Machine. New York: Plenum Press. ISBN 978-0-306-45567-4.
	Clemence, Gerald M. (1947). "The Relativity Effect in Planetary Motions". Reviews of Modern Physics. 19 (4): 361–364. Bibcode:1947RvMP...19..361C. doi:10.1103/RevModPhys.19.361.
	Gilvarry, John J. (1953). "Relativity Precession of the Asteroid Icarus". Physical Review. 89 (5): 1046. Bibcode:1953PhRv...89.1046G. doi:10.1103/PhysRev.89.1046.
	Brown, Kevin. "6.2 Anomalous Precession". Reflections on Relativity. MathPages. Retrieved May 22, 2008.
	Hall, Shannon (March 24, 2020). "Life on the Planet Mercury? 'It's Not Completely Nuts' - A new explanation for the rocky world's jumbled landscape opens a possibility that it could have had ingredients for habitability". The New York Times. Retrieved March 26, 2020.
	Rodriguez, J. Alexis P.; Leonard, Gregory J.; Kargel, Jeffrey S.; Domingue, Deborah; Berman, Daniel C.; Banks, Maria; Zarroca, Mario; Linares, Rogelio; Marchi, Simone; Baker, Victor R.; Webster, Kevin D.; Sykes, Mark (March 16, 2020). "The Chaotic Terrains of Mercury Reveal a History of Planetary Volatile Retention and Loss in the Innermost Solar System". Scientific Reports. 10 (4737): 4737. Bibcode:2020NatSR..10.4737R. doi:10.1038/s41598-020-59885-5. PMC 7075900. PMID 32179758.
	Menzel, Donald H. (1964). A Field Guide to the Stars and Planets. The Peterson Field Guide Series. Boston: Houghton Mifflin Co. pp. 292–293.
	Tezel, Tunç (January 22, 2003). "Total Solar Eclipse of 2006 March 29". Department of Physics at Fizik Bolumu in Turkey. Retrieved May 24, 2008.
	Mallama, Anthony (2011). "Planetary magnitudes". Sky and Telescope. 121 (1): 51–56.
	Espenak, Fred (1996). "NASA Reference Publication 1349; Venus: Twelve year planetary ephemeris, 1995–2006". Twelve Year Planetary Ephemeris Directory. NASA. Retrieved May 24, 2008.
	Walker, John. "Mercury Chaser's Calculator". Fourmilab Switzerland. Retrieved May 29, 2008. (look at 1964 and 2013)
	"Mercury Elongation and Distance". Archived from the original on May 11, 2013. Retrieved May 30, 2008. – Numbers generated using the Solar System Dynamics Group, Horizons On-Line Ephemeris System
	Kelly, Patrick, ed. (2007). Observer's Handbook 2007. Royal Astronomical Society of Canada. ISBN 978-0-9738109-3-6.
	Curtis, A. C. (October 1972). "Finding Venus or Mercury in daylight". Journal of the British Astronomical Association. 82: 438–439. Bibcode:1972JBAA...82..438C.
	Alers, Paul E. (March 17, 2011). "Celebrating Mercury Orbit". NASA Multimedia. Retrieved March 18, 2011.
	"NASA spacecraft now circling Mercury – a first". NBC News. March 17, 2011. Retrieved March 24, 2011.
	Baumgardner, Jeffrey; Mendillo, Michael; Wilson, Jody K. (2000). "A Digital High-Definition Imaging System for Spectral Studies of Extended Planetary Atmospheres. I. Initial Results in White Light Showing Features on the Hemisphere of Mercury Unimaged by Mariner 10". The Astronomical Journal. 119 (5): 2458–2464. Bibcode:2000AJ....119.2458B. doi:10.1086/301323.
	Schaefer, Bradley E. (2007). "The Latitude and Epoch for the Origin of the Astronomical Lore in Mul.Apin". American Astronomical Society Meeting 210, #42.05. 38: 157. Bibcode:2007AAS...210.4205S.
	Hunger, Hermann; Pingree, David (1989). "MUL.APIN: An Astronomical Compendium in Cuneiform". Archiv für Orientforschung. 24: 146.
	"MESSENGER: Mercury and Ancient Cultures". NASA JPL. 2008. Retrieved April 7, 2008.
	Dunne, James A.; Burgess, Eric (1978). "Chapter One". The Voyage of Mariner 10 – Mission to Venus and Mercury. NASA History Office.
	Στίλβων, Ἑρμῆς. Liddell, Henry George; Scott, Robert; A Greek–English Lexicon at the Perseus Project.
	"Greek Names of the Planets". April 25, 2010. Retrieved July 14, 2012. "Ermis is the Greek name of the planet Mercury, which is the closest planet to the Sun. It is named after the Greek God of commerce, Ermis or Hermes, who was also the messenger of the Ancient Greek gods." See also the Greek article about the planet.
	Antoniadi, Eugène Michel (1974). The Planet Mercury. Translated from French by Moore, Patrick. Shaldon, Devon: Keith Reid Ltd. pp. 9–11. ISBN 978-0-904094-02-2.
	Duncan, John Charles (1946). Astronomy: A Textbook. Harper & Brothers. p. 125. "The symbol for Mercury represents the Caduceus, a wand with two serpents twined around it, which was carried by the messenger of the gods."
	Heath, Sir Thomas (1921). A History of Greek Mathematics. II. Oxford: Clarendon Press. pp. vii, 273.
	Goldstein, Bernard R. (1996). "The Pre-telescopic Treatment of the Phases and Apparent Size of Venus". Journal for the History of Astronomy. 27: 1. Bibcode:1996JHA....27....1G. doi:10.1177/002182869602700101. S2CID 117218196.
	Kelley, David H.; Milone, E. F.; Aveni, Anthony F. (2004). Exploring Ancient Skies: An Encyclopedic Survey of Archaeoastronomy. Birkhäuser. ISBN 978-0-387-95310-6.
	De Groot, Jan Jakob Maria (1912). Religion in China: universism. a key to the study of Taoism and Confucianism. American lectures on the history of religions. 10. G. P. Putnam's Sons. p. 300. Retrieved January 8, 2010.
	Crump, Thomas (1992). The Japanese numbers game: the use and understanding of numbers in modern Japan. Nissan Institute/Routledge Japanese studies series. Routledge. pp. 39–40. ISBN 978-0-415-05609-0.
	Hulbert, Homer Bezaleel (1909). The passing of Korea. Doubleday, Page & company. p. 426. Retrieved January 8, 2010.
	Pujari, R.M.; Kolhe, Pradeep; Kumar, N. R. (2006). Pride of India: A Glimpse Into India's Scientific Heritage. Samskrita Bharati. ISBN 978-81-87276-27-2.
	Bakich, Michael E. (2000). The Cambridge Planetary Handbook. Cambridge University Press. ISBN 978-0-521-63280-5.
	Milbrath, Susan (1999). Star Gods of the Maya: Astronomy in Art, Folklore and Calendars. University of Texas Press. ISBN 978-0-292-75226-9.
	Samsó, Julio; Mielgo, Honorino (1994). "Ibn al-Zarqālluh on Mercury". Journal for the History of Astronomy. 25 (4): 289–96 [292]. Bibcode:1994JHA....25..289S. doi:10.1177/002182869402500403. S2CID 118108131.
	Hartner, Willy (1955). "The Mercury Horoscope of Marcantonio Michiel of Venice". Vistas in Astronomy. 1 (1): 84–138. Bibcode:1955VA......1...84H. doi:10.1016/0083-6656(55)90016-7. at pp. 118–122.
	Ansari, S. M. Razaullah (2002). History of oriental astronomy: proceedings of the joint discussion-17 at the 23rd General Assembly of the International Astronomical Union, organised by the Commission 41 (History of Astronomy), held in Kyoto, August 25–26, 1997. Springer Science+Business Media. p. 137. ISBN 1-4020-0657-8.
	Goldstein, Bernard R. (1969). "Some Medieval Reports of Venus and Mercury Transits". Centaurus. 14 (1): 49–59. Bibcode:1969Cent...14...49G. doi:10.1111/j.1600-0498.1969.tb00135.x.
	Ramasubramanian, K.; Srinivas, M. S.; Sriram, M. S. (1994). "Modification of the Earlier Indian Planetary Theory by the Kerala Astronomers (c. 1500 AD) and the Implied Heliocentric Picture of Planetary Motion" (PDF). Current Science. 66: 784–790. Archived from the original (PDF) on December 23, 2010. Retrieved April 23, 2010.
	Sinnott, Roger W.; Meeus, Jean (1986). "John Bevis and a Rare Occultation". Sky and Telescope. 72: 220. Bibcode:1986S&T....72..220S.
	Ferris, Timothy (2003). Seeing in the Dark: How Amateur Astronomers. Simon and Schuster. ISBN 978-0-684-86580-5.
	Colombo, Giuseppe; Shapiro, Irwin I. (November 1965). "The Rotation of the Planet Mercury". SAO Special Report #188R. 188: 188. Bibcode:1965SAOSR.188.....C.
	Holden, Edward S. (1890). "Announcement of the Discovery of the Rotation Period of Mercury [by Professor Schiaparelli]". Publications of the Astronomical Society of the Pacific. 2 (7): 79. Bibcode:1890PASP....2...79H. doi:10.1086/120099.
	Davies, Merton E.; Dwornik, Stephen E.; Gault, Donald E.; Strom, Robert G. (1978). "Surface Mapping". Atlas of Mercury. NASA Office of Space Sciences. Retrieved May 28, 2008.
	Evans, John V.; Brockelman, Richard A.; Henry, John C.; Hyde, Gerald M.; Kraft, Leon G.; Reid, Wyatt A.; Smith, W. W. (1965). "Radio Echo Observations of Venus and Mercury at 23 cm Wavelength". Astronomical Journal. 70: 487–500. Bibcode:1965AJ.....70..486E. doi:10.1086/109772.
	Moore, Patrick (2000). The Data Book of Astronomy. New York: CRC Press. p. 483. ISBN 978-0-7503-0620-1.
	Butrica, Andrew J. (1996). "Chapter 5". To See the Unseen: A History of Planetary Radar Astronomy. NASA History Office, Washington D.C. ISBN 978-0-16-048578-7.
	Pettengill, Gordon H.; Dyce, Rolf B. (1965). "A Radar Determination of the Rotation of the Planet Mercury". Nature. 206 (1240): 451–2. Bibcode:1965Natur.206Q1240P. doi:10.1038/2061240a0. S2CID 31525579.
	"Mercury". Eric Weisstein's World of Astronomy. Wolfram Research. Retrieved April 18, 2021.
	Murray, Bruce C.; Burgess, Eric (1977). Flight to Mercury. Columbia University Press. ISBN 978-0-231-03996-3.
	Colombo, Giuseppe (1965). "Rotational Period of the Planet Mercury". Nature. 208 (5010): 575. Bibcode:1965Natur.208..575C. doi:10.1038/208575a0. S2CID 4213296.
	Davies, Merton E.; et al. (1976). "Mariner 10 Mission and Spacecraft". SP-423 Atlas of Mercury. NASA JPL. Retrieved April 7, 2008.
	Golden, Leslie M. (1977). A Microwave Interferometric Study of the Subsurface of the Planet Mercury (Thesis). University of California, Berkeley. Bibcode:1977PhDT.........9G.
	Mitchell, David L.; De Pater, Imke (1994). "Microwave Imaging of Mercury's Thermal Emission at Wavelengths from 0.3 to 20.5 cm (1994)". Icarus. 110 (1): 2–32. Bibcode:1994Icar..110....2M. doi:10.1006/icar.1994.1105.
	Dantowitz, Ronald F.; Teare, Scott W.; Kozubal, Marek J. (2000). "Ground-based High-Resolution Imaging of Mercury". Astronomical Journal. 119 (4): 2455–2457. Bibcode:2000AJ....119.2455D. doi:10.1086/301328.
	Harmon, John K.; Slade, Martin A.; Butler, Bryan J.; Head III, James W.; Rice, Melissa S.; Campbell, Donald B. (2007). "Mercury: Radar images of the equatorial and midlatitude zones". Icarus. 187 (2): 374–405. Bibcode:2007Icar..187..374H. doi:10.1016/j.icarus.2006.09.026.
	Webster, Guy (June 10, 2014). "Mercury Passes in Front of the Sun, as Seen From Mars". NASA. Retrieved June 10, 2014.
	Zacny, Kris (July 2, 2015). Inner Solar System: Prospective Energy and Material Resources. Springer International Publishing. p. 154. ISBN 9783319195698.
	Wagner, Sam; Wie, Bong (November 2015). "Hybrid Algorithm for Multiple Gravity-Assist and Impulsive Delta-V Maneuvers". Journal of Guidance, Control, and Dynamics. 38 (11): 2096–2107. Bibcode:2015JGCD...38.2096W. doi:10.2514/1.G000874.
	"Mercury" (PDF). NASA Jet Propulsion Laboratory. May 5, 2008. Retrieved April 26, 2021.
	Leipold, Manfred E.; Seboldt, W.; Lingner, Stephan; Borg, Erik; Herrmann, Axel Siegfried; Pabsch, Arno; Wagner, O.; Brückner, Johannes (1996). "Mercury sun-synchronous polar orbiter with a solar sail". Acta Astronautica. 39 (1): 143–151. Bibcode:1996AcAau..39..143L. doi:10.1016/S0094-5765(96)00131-2.
	Dunne, James A. & Burgess, Eric (1978). "Chapter Four". The Voyage of Mariner 10 – Mission to Venus and Mercury. NASA History Office. Retrieved May 28, 2008.
	Phillips, Tony (October 1976). "NASA 2006 Transit of Mercury". SP-423 Atlas of Mercury. NASA. Retrieved April 7, 2008.
	"BepiColumbo – Background Science". European Space Agency. Retrieved June 18, 2017.
	Malik, Tariq (August 16, 2004). "MESSENGER to test theory of shrinking Mercury". USA Today. Retrieved May 23, 2008.
	Davies ME, et al. (1978). "Mariner 10 Mission and Spacecraft". Atlas of Mercury. NASA Office of Space Sciences. Retrieved May 30, 2008.
	Ness, Norman F. (1978). "Mercury – Magnetic field and interior". Space Science Reviews. 21 (5): 527–553. Bibcode:1978SSRv...21..527N. doi:10.1007/BF00240907. S2CID 120025983.
	Aharonson, Oded; Zuber, Maria T; Solomon, Sean C (2004). "Crustal remanence in an internally magnetized non-uniform shell: a possible source for Mercury's magnetic field?". Earth and Planetary Science Letters. 218 (3–4): 261–268. Bibcode:2004E&PSL.218..261A. doi:10.1016/S0012-821X(03)00682-4.
	Dunne, James A. & Burgess, Eric (1978). "Chapter Eight". The Voyage of Mariner 10 – Mission to Venus and Mercury. NASA History Office.
	Grayzeck, Ed (April 2, 2008). "Mariner 10". NSSDC Master Catalog. NASA. Retrieved April 7, 2008.
	"MESSENGER Engine Burn Puts Spacecraft on Track for Venus". SpaceRef.com. 2005. Retrieved March 2, 2006.
	"Countdown to MESSENGER's Closest Approach with Mercury". Johns Hopkins University Applied Physics Laboratory. January 14, 2008. Archived from the original on May 13, 2013. Retrieved May 30, 2008.
	"MESSENGER Gains Critical Gravity Assist for Mercury Orbital Observations". MESSENGER Mission News. September 30, 2009. Archived from the original on May 10, 2013. Retrieved September 30, 2009.
	"NASA extends spacecraft's Mercury mission". UPI. November 15, 2011. Retrieved November 16, 2011.
	"MESSENGER: Fact Sheet" (PDF). Applied Physics Laboratory. February 2011. Retrieved August 21, 2017.
	Wall, Mike (March 29, 2015). "NASA Mercury Probe Trying to Survive for Another Month". Space.com. Retrieved April 4, 2015.
	Chang, Kenneth (April 27, 2015). "NASA's Messenger Mission Is Set to Crash Into Mercury". The New York Times. Retrieved April 27, 2015.
	Corum, Jonathan (April 30, 2015). "Messenger's Collision Course With Mercury". The New York Times. Retrieved April 30, 2015.
	"Best Determination of MESSENGER's Impact Location". MESSENGER Featured Images. Johns Hopkins – Applied Physics Lab. June 3, 2015. Retrieved April 29, 2015.
	"ESA gives go-ahead to build BepiColombo". European Space Agency. February 26, 2007. Retrieved May 29, 2008.
	"BepiColombo Fact Sheet". European Space Agency. December 1, 2016. Retrieved December 19, 2016.
	
		"Objectives". European Space Agency. February 21, 2006. Retrieved May 29, 2008.
	
	External links
	Listen to this article (41 minutes)
	Menu
	0:00
	Spoken Wikipedia icon
	This audio file was created from a revision of this article dated 16 January 2008, and does not reflect subsequent edits.
	(Audio help · More spoken articles)
	Mercury
	at Wikipedia's sister projects
	
		Definitions from Wiktionary
		Media from Wikimedia Commons
		Quotations from Wikiquote
		Texts from Wikisource
		Textbooks from Wikibooks
		Resources from Wikiversity
		Data from Wikidata
	
		Atlas of Mercury. NASA. 1978. SP-423.
		Mercury nomenclature and map with feature names from the USGS/IAU Gazetteer of Planetary Nomenclature
		Equirectangular map of Mercury by Applied Coherent Technology Corp
		3D globe of Mercury by Google
		Mercury at Solarviews.com
		Mercury by Astronomy Cast
		MESSENGER mission web site
		BepiColombo mission web site
	
		vte
	
	Mercury
	
		Outline
	
	Geography	
	General	
	
		Albedo features Atmosphere Features Magnetic field
	
	Regions	
	
		Quadrangles
	
	Quadrangles	
	
		Bach Beethoven Borealis Debussy Derain Discovery Eminescu Hokusai Kuiper Michelangelo Neruda Shakespeare Raditladi Tolstoj Victoria
	
	Mountains and
	volcanoes	
	
		Caloris Montes
	
	Plains and
	plateaus	
	
		Borealis Planitia Budh Planitia Caloris Planitia Odin Planitia Sobkou Planitia Suisei Planitia Tir Planitia
	
	Canyons and
	valleys	
	
		Arecibo Catena Arecibo Vallis Goldstone Catena Goldstone Vallis Haystack Catena Haystack Vallis Simeiz Vallis
	
	Ridges and
	rupes	
	
		Antoniadi Dorsum Adventure Rupes Beagle Rupes Discovery Rupes Enterprise Rupes Hero Rupes Resolution Rupes Santa María Rupes Victoria Rupes
	
	Basins and
	fossae	
	
		Caloris Basin Pantheon Fossae Raditladi Basin Rembrandt Basin Skinakas (hypothetical basin)
	
	Craters	
	
		Abedin Abu Nuwas Africanus Horton Ahmad Baba Ailey Aksakov Akutagawa Al-Akhtal Alencar Al-Hamadhani Al-Jāhiz Alver Amaral Amru Al-Qays Andal Aneirin Angelou Anguissola Anyte Apollodorus Aristoxenes Aśvaghosa Atget Bach Balagtas Balanchine Baranauskas Balzac Bartók Barma Bashō Beckett Beethoven Bek Belinskij Bello Benoit Berkel Bernini Bjornson Boccaccio Boethius Botticelli Brahms Bramante Brontë Bruegel Brunelleschi Burns Byron Callicrates Camoes Carducci Carolan Calvino Cervantes Cézanne Chaikovskij Chao Meng-Fu Chekov Chiang K'ui Chong Ch'ol Chopin Chu Ta Coleridge Copland Copley Couperin Cunningham Dali Dario De Graft Debussy Degas Delacroix Derain Derzhavin Desprez Dickens Dominici Donne Dostoevskij Dowland Durer Dvorak Eastman Eitoku Eminescu Enheduanna Enwonwu Equiano Faulkner Fet Firdousi Flaubert Flaiano Futabatei Gainsborough Gauguin Geddes Ghiberti Gibran Giotto Glinka Gluck Goethe Gogol Goya Grieg Guido d'Arezzo Hals Han Kan Handel Harunobu Hauptmann Hawthorne Haydn Heine Hemingway Hesiod Hiroshige Hitomaro Hodgkins Hokusai Holbein Holberg Holst Homer Horace Hovnatanian Hugo Hun Kal Ibsen Ictinus Imhotep Ives Izquierdo Janáček Jokai Judah Ha-Levi Kalidasa Karsh Keats Kenko Kertész Khansa Kipling Kōshō Kuan Han-ch'ing Kuiper Kulthum Kunisada Kurosawa Lange Larrocha Leopardi Lermontov Lessing Li Ch'ing-Chao Li Po Liang K'ai Liszt Lu Hsun Lysippus Ma Chih-Yuan Machaut Mahler Mansart Mansur March Mark Twain Martí Martial Matabei Matisse Melville Mena Mendes Pinto Michelangelo Mickiewicz Milton Mistral Mofolo Molière Monet Monteverdi Moody Mozart Munch Munkácsy Murasaki Mussorgskij Myron Nabokov Nampeyo Navoi Nawahi Neruda Nureyev Nervo Neumann Nizami Okyo Oskison Ovid Petrarch Phidias Picasso Poe Polygnotus Praxiteles Prokofiev Qi Baishi Rachmaninoff Raden Saleh Raditladi Rameau Raphael Rembrandt Renoir Rivera Rodin Rudaki Sander Scarlatti Schubert Shakespeare Sholem Aleichem Sinan Stravinsky Sullivan Sveinsdóttir Titian To Ngoc Van Tolstoj Verdi Villa-Lobos Vivaldi Vyasa Xiao Zhao Yeats Zola
	
	Other	
	
		Geology Ghost craters Inter-crater plains
	
	Moons	
	
		Hypothetical moon of Mercury
	
	Astronomy	
	Transits	
	
		Earth Mars
	
	Asteroids	
	
		Mercury-crossing asteroids
	
	Exploration	
	Current
	and past	
	
		Mariner 10 (1973–1975) MESSENGER (2004–2015) BepiColombo (in transit)
	
	Proposed	
	
		Mercury-P (~2031)
	
	See also	
	
		Colonization of Mercury
	
	Related	
	
		Fiction Sub-Earth
	
		Category Category Portal Portal
	
		vte
	
	Solar System
	
		Stars portalSpaceflight portalOuter space portal
	
	Authority control Edit this at Wikidata
	Categories:
	
		Solar SystemMercury (planet)Planets of the Solar SystemTerrestrial planetsAstronomical objects known since antiquity
	
	Navigation menu
	
		Not logged in
		Talk
		Contributions
		Create account
		Log in
	
		Article
		Talk
	
		Read
		View source
		View history
	
	Search
	
		Main page
		Contents
		Current events
		Random article
		About Wikipedia
		Contact us
		Donate
	
	Contribute
	
		Help
		Learn to edit
		Community portal
		Recent changes
		Upload file
	
	Tools
	
		What links here
		Related changes
		Special pages
		Permanent link
		Page information
		Cite this page
		Wikidata item
	
	Print/export
	
		Download as PDF
		Printable version
	
	In other projects
	
		Wikimedia Commons
		Wikinews
		Wikiquote
		Wikiversity
	
	Languages
	
		Башҡортса
		Чӑвашла
		Հայերեն
		日本語
		Нохчийн
		Русский
		Саха тыла
		Татарча/tatarça
		Удмурт
	
	Edit links
	
		This page was last edited on 26 August 2021, at 23:21 (UTC).
		Text is available under the Creative Commons Attribution-ShareAlike License; additional terms may apply. By using this site, you agree to the Terms of Use and Privacy Policy. Wikipedia® is a registered trademark of the Wikimedia Foundation, Inc., a non-profit organization.
	
		Privacy policy
		About Wikipedia
		Disclaimers
		Contact Wikipedia
		Mobile view
		Developers
		Statistics
		Cookie statement
	
		Wikimedia Foundation
		Powered by MediaWiki
	
	`

	dutch = `

	Wiki Loves Monuments: your chance to support Russian cultural heritage! Photograph a monument and win!
	Featured article
	Page semi-protected
	Moon
	From Wikipedia, the free encyclopedia
	Jump to navigation
	Jump to search
	This article is about Earth's natural satellite. For moons in general, see Natural satellite. For other uses, see Moon (disambiguation).
	Moon Moon symbolFull Moon in the darkness of the night sky. It is patterned with a mix of light-tone regions and darker, irregular blotches, and scattered with varying sizes of impact craters, circles surrounded by out-thrown rays of bright ejecta.
	The Moon with its fully illuminated near side (north is at top)
	Designations
	Designation
		Earth I
	Alternative names
		
	
		LunaSelene (poetic)Cynthia (poetic)
	
	Adjectives	
	
		LunarSelenian (poetic)Cynthian (poetic)Moonly (poetic)
	
	Orbital characteristics
	Epoch J2000
	Perigee	362600 km
	(356400–370400 km)
	Apogee	405400 km
	(404000–406700 km)
	Semi-major axis
		384399 km  (1.28 ls, 0.00257 AU)[1]
	Eccentricity	0.0549[1]
	Orbital period
		27.321661 d
	(27 d 7 h 43 min 11.5 s[1])
	Synodic period
		29.530589 d
	(29 d 12 h 44 min 2.9 s)
	Average orbital speed
		1.022 km/s
	Inclination	5.145° to the ecliptic[2][a]
	Longitude of ascending node
		Regressing by one revolution in 18.61 years
	Argument of perigee
		Progressing by one
	revolution in 8.85 years
	Satellite of	Earth[b][3]
	Physical characteristics
	Mean radius
		1737.4 km  
	(0.2727 of Earth's)[1][4][5]
	Equatorial radius
		1738.1 km  
	(0.2725 of Earth's)[4]
	Polar radius
		1736.0 km  
	(0.2731 of Earth's)[4]
	Flattening	0.0012[4]
	Circumference	10921 km  (equatorial)
	Surface area
		3.793×107 km2  
	(0.074 of Earth's)
	Volume	2.1958×1010 km3  
	(0.020 of Earth's)[4]
	Mass	7.342×1022 kg  
	(0.012300 of Earth's)[1][4][6]
	Mean density
		3.344 g/cm3[1][4]
	0.606 × Earth
	Surface gravity
		1.62 m/s2  (0.1654 g)[4]
	Moment of inertia factor
		0.3929±0.0009[7]
	Escape velocity
		2.38 km/s
	(8600 km/h; 5300 mph)
	Rotation period
		29.530589 d
	(29 d 12 h 44 min 2.9 s; synodic; solar day) (spin-orbit locked)
	Sidereal rotation period
		27.321661 d  (spin-orbit locked)
	Equatorial rotation velocity
		4.627 m/s
	Axial tilt
		
	
		1.5424° to ecliptic[8]
		6.687° to orbit plane[2]
		24° to Earth's equator [9]
	
	North pole right ascension
		
	
		17h 47m 26s
		266.86°[10]
	
	North pole declination
		65.64°[10]
	Albedo	0.136[11]
	Surface temp. 	min 	mean 	max
	Equator 	100 K[12] 	250 K 	390 K[12]
	85°N  		150 K 	230 K[13]
	Apparent magnitude
		
	
		−2.5 to −12.9[c]
		−12.74  (mean full moon)[4]
	
	Angular diameter
		29.3 to 34.1 arcminutes[4][d]
	Atmosphere[14]
	Surface pressure
		
	
		10−7 Pa (1 picobar)  (day)
		10−10 Pa (1 femtobar)   
		(night)[e]
	
	Composition by volume	
	
		HeArNeNaKHRn
	
	The Moon is Earth's only natural satellite. At about one-quarter the diameter of Earth (comparable to the width of Australia),[15] it is the largest natural satellite in the Solar System relative to the size of its planet,[f] the fifth largest satellite in the Solar System overall, and is larger than any known dwarf planet. Orbiting Earth at an average distance of 384,400 km (238,900 mi),[16] or about 30 times Earth's diameter, its gravitational influence slightly lengthens Earth's day and is the main driver of Earth's tides. The Moon is classified as a planetary-mass object and a differentiated rocky body, and lacks any significant atmosphere, hydrosphere, or magnetic field. Its surface gravity is about one-sixth of Earth's (0.1654 g); Jupiter's moon Io is the only satellite in the Solar System known to have a higher surface gravity and density.
	
	The Moon's orbit around Earth has a sidereal period of 27.3 days. During each synodic period of 29.5 days, the amount of visible surface illuminated by the Sun varies from none up to 100%, resulting in lunar phases that form the basis for the months of a lunar calendar. The Moon is tidally locked to Earth, which means that the length of a full rotation of the Moon on its own axis causes its same side (the near side) to always face Earth, and the somewhat longer lunar day is the same as the synodic period. That said, 59% of the total lunar surface can be seen from Earth through shifts in perspective due to libration.[17]
	
	The most widely accepted origin explanation posits that the Moon formed about 4.51 billion years ago, not long after Earth, out of the debris from a giant impact between the planet and a hypothesized Mars-sized body called Theia. It then receded to a wider orbit because of tidal interaction with the Earth. The near side of the Moon is marked by dark volcanic maria ("seas"), which fill the spaces between bright ancient crustal highlands and prominent impact craters. Most of the large impact basins and mare surfaces were in place by the end of the Imbrian period, some three billion years ago. The lunar surface is relatively non-reflective, with a reflectance just slightly brighter than that of worn asphalt. However, because it has a large angular diameter, the full moon is the brightest celestial object in the night sky. The Moon's apparent size is nearly the same as that of the Sun, allowing it to cover the Sun almost completely during a total solar eclipse.
	
	Both the Moon's prominence in the earthly sky and its regular cycle of phases have provided cultural references and influences for human societies throughout history. Such influences can be found in language, calendar systems, art, and mythology. The first artificial object to reach the Moon was the Soviet Union's Luna 2 uncrewed spacecraft in 1959; this was followed by the first successful soft landing by Luna 9 in 1966. The only human lunar missions to date have been those of the United States' Apollo program, which landed twelve men on the surface between 1969 and 1972. These and later uncrewed missions returned lunar rocks that have been used to develop a detailed geological understanding of the Moon's origins, internal structure, and subsequent history.
	Contents
	
		1 Name and etymology
		2 Formation
		3 Physical characteristics
			3.1 Internal structure
				3.1.1 Magnetic field
			3.2 Surface geology
				3.2.1 Volcanic features
				3.2.2 Impact craters
				3.2.3 Gravitational field
				3.2.4 Lunar swirls
				3.2.5 Presence of water
			3.3 Surface conditions
				3.3.1 Atmosphere
				3.3.2 Dust
		4 Earth–Moon system
			4.1 Lunar distance
			4.2 Orbit
			4.3 Relative size
			4.4 Appearance from Earth
				4.4.1 Eclipses
			4.5 Tidal effects
		5 Observation and exploration
			5.1 Before spaceflight
			5.2 1959–1970s
				5.2.1 Soviet missions
				5.2.2 United States missions
			5.3 1970s – present
			5.4 Future
				5.4.1 Planned commercial missions
		6 Human presence
			6.1 Human impact
			6.2 Infrastructure
			6.3 Astronomy from the Moon
			6.4 Living on the Moon
		7 Legal status
			7.1 Coordination
		8 In culture and life
			8.1 Mythology
			8.2 Calendar
			8.3 Lunar effect
		9 See also
		10 Notes
		11 References
			11.1 Citations
		12 Further reading
		13 External links
			13.1 Cartographic resources
			13.2 Observation tools
			13.3 General
	
	Name and etymology
	See also: List of lunar deities
	
	The usual English proper name for Earth's natural satellite is simply the Moon, with a capital M.[18][19] The noun moon is derived from Old English mōna, which (like all its Germanic cognates) stems from Proto-Germanic *mēnōn,[20] which in turn comes from Proto-Indo-European *mēnsis "month"[21] (from earlier *mēnōt, genitive *mēneses) which may be related to the verb "measure" (of time).[22]
	The Moon is prominently featured in Vincent van Gogh's 1889 painting, The Starry Night
	
	Occasionally, the name Luna /ˈluːnə/ is used in scientific writing[23] and especially in science fiction to distinguish the Earth's moon from others, while in poetry "Luna" has been used to denote personification of the Moon.[24] Cynthia /ˈsɪnθiə/ is another poetic name, though rare, for the Moon personified as a goddess,[25] while Selene /səˈliːniː/ (literally "Moon") is the Greek goddess of the Moon.
	
	The usual English adjective pertaining to the Moon is "lunar", derived from the Latin word for the Moon, lūna. The adjective selenian /səliːniən/,[26] derived from the Greek word for the Moon, σελήνη selēnē, and used to describe the Moon as a world rather than as an object in the sky, is rare,[27] while its cognate selenic was originally a rare synonym[28] but now nearly always refers to the chemical element selenium.[29] The Greek word for the Moon does however provide us with the prefix seleno-, as in selenography, the study of the physical features of the Moon, as well as the element name selenium.[30][31]
	
	The Greek goddess of the wilderness and the hunt, Artemis, equated with the Roman Diana, one of whose symbols was the Moon and who was often regarded as the goddess of the Moon, was also called Cynthia, from her legendary birthplace on Mount Cynthus.[32] These names – Luna, Cynthia and Selene – are reflected in technical terms for lunar orbits such as apolune, pericynthion and selenocentric.
	The Moon
	Near side of the Moon
	Far side of the Moon
	Lunar north pole
	Lunar south pole
	Formation
	Main articles: Origin of the Moon, Giant-impact hypothesis, and Circumplanetary disk
	
	Isotope dating of lunar samples suggests the Moon formed around 50 million years after the origin of the Solar System.[33][34] Historically, several formation mechanisms have been proposed,[35] but none satisfactorily explained the features of the Earth–Moon system. A fission of the Moon from Earth's crust through centrifugal force[36] would require too great an initial rotation rate of Earth.[37] Gravitational capture of a pre-formed Moon[38] depends on an unfeasibly extended atmosphere of Earth to dissipate the energy of the passing Moon.[37] A co-formation of Earth and the Moon together in the primordial accretion disk does not explain the depletion of metals in the Moon.[37] None of these hypotheses can account for the high angular momentum of the Earth–Moon system.[39]
	File:Evolution of the Moon.ogvPlay media
	The evolution of the Moon and a tour of the Moon
	
	The prevailing theory is that the Earth–Moon system formed after a giant impact of a Mars-sized body (named Theia) with the proto-Earth. The impact blasted material into Earth's orbit and then the material accreted and formed the Moon[40][41] just beyond the Earth's Roche limit of ~2.56 R⊕.[42] This theory best explains the evidence.
	
	Giant impacts are thought to have been common in the early Solar System. Computer simulations of giant impacts have produced results that are consistent with the mass of the lunar core and the angular momentum of the Earth–Moon system. These simulations also show that most of the Moon derived from the impactor, rather than the proto-Earth.[43] However, more recent simulations suggest a larger fraction of the Moon derived from the proto-Earth.[44][45][46][47] Other bodies of the inner Solar System such as Mars and Vesta have, according to meteorites from them, very different oxygen and tungsten isotopic compositions compared to Earth. However, Earth and the Moon have nearly identical isotopic compositions. The isotopic equalization of the Earth-Moon system might be explained by the post-impact mixing of the vaporized material that formed the two,[48] although this is debated.[49]
	
	The impact released energy and then the released material re-accreted into the Earth–Moon system. This would have melted the outer shell of Earth, and thus formed a magma ocean.[50][51] Similarly, the newly formed Moon would also have been affected and had its own lunar magma ocean; its depth is estimated from about 500 km (300 miles) to 1,737 km (1,079 miles).[50]
	Oceanus Procellarum ("Ocean of Storms")
	Ancient rift valleys – rectangular structure (visible – topography – GRAIL gravity gradients)
	Ancient rift valleys – context
	Ancient rift valleys – closeup (artist's concept)
	
	While the giant-impact theory explains many lines of evidence, some questions are still unresolved, most of which involve the Moon's composition.[52]
	
	In 2001, a team at the Carnegie Institute of Washington reported the most precise measurement of the isotopic signatures of lunar rocks.[53] The rocks from the Apollo program had the same isotopic signature as rocks from Earth, differing from almost all other bodies in the Solar System. This observation was unexpected, because most of the material that formed the Moon was thought to come from Theia and it was announced in 2007 that there was less than a 1% chance that Theia and Earth had identical isotopic signatures.[54] Other Apollo lunar samples had in 2012 the same titanium isotopes composition as Earth,[55] which conflicts with what is expected if the Moon formed far from Earth or is derived from Theia. These discrepancies may be explained by variations of the giant-impact theory.
	Physical characteristics
	
	The Moon is a very slightly scalene ellipsoid due to tidal stretching, with its long axis displaced 30° from facing the Earth, due to gravitational anomalies from impact basins. Its shape is more elongated than current tidal forces can account for. This 'fossil bulge' indicates that the Moon solidified when it orbited at half its current distance to the Earth, and that it is now too cold for its shape to adjust to its orbit.[56]
	Internal structure
	Main article: Internal structure of the Moon
	Lunar surface chemical composition[57] Compound 	Formula 	Composition
	Maria 	Highlands
	silica 	SiO2 	45.4% 	45.5%
	alumina 	Al2O3 	14.9% 	24.0%
	lime 	CaO 	11.8% 	15.9%
	iron(II) oxide 	FeO 	14.1% 	5.9%
	magnesia 	MgO 	9.2% 	7.5%
	titanium dioxide 	TiO2 	3.9% 	0.6%
	sodium oxide 	Na2O 	0.6% 	0.6%
		  99.9% 	100.0%
	
	The Moon is a differentiated body that was initially in hydrostatic equilibrium but has since departed from this condition.[58] It has a geochemically distinct crust, mantle, and core. The Moon has a solid iron-rich inner core with a radius possibly as small as 240 kilometres (150 mi) and a fluid outer core primarily made of liquid iron with a radius of roughly 300 kilometres (190 mi). Around the core is a partially molten boundary layer with a radius of about 500 kilometres (310 mi).[59][60] This structure is thought to have developed through the fractional crystallization of a global magma ocean shortly after the Moon's formation 4.5 billion years ago.[61]
	
	Crystallization of this magma ocean would have created a mafic mantle from the precipitation and sinking of the minerals olivine, clinopyroxene, and orthopyroxene; after about three-quarters of the magma ocean had crystallised, lower-density plagioclase minerals could form and float into a crust atop.[62] The final liquids to crystallise would have been initially sandwiched between the crust and mantle, with a high abundance of incompatible and heat-producing elements.[1] Consistent with this perspective, geochemical mapping made from orbit suggests a crust of mostly anorthosite.[14] The Moon rock samples of the flood lavas that erupted onto the surface from partial melting in the mantle confirm the mafic mantle composition, which is more iron-rich than that of Earth.[1] The crust is on average about 50 kilometres (31 mi) thick.[1]
	
	The Moon is the second-densest satellite in the Solar System, after Io.[63] However, the inner core of the Moon is small, with a radius of about 350 kilometres (220 mi) or less,[1] around 20% of the radius of the Moon. Its composition is not well understood, but is probably metallic iron alloyed with a small amount of sulfur and nickel; analyses of the Moon's time-variable rotation suggest that it is at least partly molten.[64] The pressure at the lunar core is estimated to be 5 GPa.[65]
	Magnetic field
	Main article: Magnetic field of the Moon
	
	The Moon has an external magnetic field of generally less than 0.2 nanoteslas,[66] or less than one hundred thousandth that of Earth. The Moon does not currently have a global dipolar magnetic field and only has crustal magnetization likely acquired early in its history when a dynamo was still operating.[67][68] However, early in its history, 4 billion years ago, its magnetic field strength was likely close to that of Earth today.[66] This early dynamo field apparently expired by about one billion years ago, after the lunar core had completely crystallized.[66] Theoretically, some of the remnant magnetization may originate from transient magnetic fields generated during large impacts through the expansion of plasma clouds. These clouds are generated during large impacts in an ambient magnetic field. This is supported by the location of the largest crustal magnetizations situated near the antipodes of the giant impact basins.[69]
	Surface geology
	Main articles: Topography of the Moon, Geology of the Moon, Moon rock, and List of lunar features
	Geological features of the Moon (near side / north pole at left, far side / south pole at right)
	Topography of the Moon measured from the Lunar Orbiter Laser Altimeter on the mission Lunar Reconnaissance Orbiter, referenced to a sphere of radius 1737.4 km
	Topography of the Moon
	
	The topography of the Moon has been measured with laser altimetry and stereo image analysis.[70] Its most extensive topographic feature is the giant far-side South Pole–Aitken basin, some 2,240 km (1,390 mi) in diameter, the largest crater on the Moon and the second-largest confirmed impact crater in the Solar System.[71][72] At 13 km (8.1 mi) deep, its floor is the lowest point on the surface of the Moon.[71][73] The highest elevations of the Moon's surface are located directly to the northeast, which might have been thickened by the oblique formation impact of the South Pole–Aitken basin.[74] Other large impact basins such as Imbrium, Serenitatis, Crisium, Smythii, and Orientale possess regionally low elevations and elevated rims.[71] The far side of the lunar surface is on average about 1.9 km (1.2 mi) higher than that of the near side.[1]
	
	The discovery of fault scarp cliffs suggest that the Moon has shrunk by about 90 metres (300 ft) within the past billion years.[75] Similar shrinkage features exist on Mercury. Mare Frigoris, a basin near the north pole long assumed to be geologically dead, has cracked and shifted. Since the Moon doesn't have tectonic plates, its tectonic activity is slow and cracks develop as it loses heat.[76]
	Volcanic features
	Main article: Volcanism on the Moon
	
	The dark and relatively featureless lunar plains, clearly seen with the naked eye, are called maria (Latin for "seas"; singular mare), as they were once believed to be filled with water;[77] they are now known to be vast solidified pools of ancient basaltic lava. Although similar to terrestrial basalts, lunar basalts have more iron and no minerals altered by water.[78] The majority of these lava deposits erupted or flowed into the depressions associated with impact basins. Several geologic provinces containing shield volcanoes and volcanic domes are found within the near side "maria".[79]
	Evidence of young lunar volcanism
	
	Almost all maria are on the near side of the Moon, and cover 31% of the surface of the near side[80] compared with 2% of the far side.[81] This is likely due to a concentration of heat-producing elements under the crust on the near side, which would have caused the underlying mantle to heat up, partially melt, rise to the surface and erupt.[62][82][83] Most of the Moon's mare basalts erupted during the Imbrian period, 3.0–3.5 billion years ago, although some radiometrically dated samples are as old as 4.2 billion years.[84] As of 2003, crater counting studies of the youngest eruptions appeared to suggest they formed no earlier than 1.2 billion years ago.[85]
	
	In 2006, a study of Ina, a tiny depression in Lacus Felicitatis, found jagged, relatively dust-free features that, because of the lack of erosion by infalling debris, appeared to be only 2 million years old.[86] Moonquakes and releases of gas also indicate some continued lunar activity.[86] Evidence of recent lunar volcanism has been identified at 70 irregular mare patches, some less than 50 million years old. This raises the possibility of a much warmer lunar mantle than previously believed, at least on the near side where the deep crust is substantially warmer because of the greater concentration of radioactive elements.[87][88][89][90] Evidence has been found for 2–10 million years old basaltic volcanism within the crater Lowell,[91][92] inside the Orientale basin. Some combination of an initially hotter mantle and local enrichment of heat-producing elements in the mantle could be responsible for prolonged activities on the far side in the Orientale basin.[93][94]
	
	The lighter-colored regions of the Moon are called terrae, or more commonly highlands, because they are higher than most maria. They have been radiometrically dated to having formed 4.4 billion years ago, and may represent plagioclase cumulates of the lunar magma ocean.[84][85] In contrast to Earth, no major lunar mountains are believed to have formed as a result of tectonic events.[95]
	
	The concentration of maria on the near side likely reflects the substantially thicker crust of the highlands of the Far Side, which may have formed in a slow-velocity impact of a second moon of Earth a few tens of millions of years after the Moon's formation.[96][97] Alternatively, it may be a consequence of asymmetrical tidal heating when the Moon was much closer to the Earth.[98]
	Impact craters
	Further information: List of craters on the Moon
	A gray, many-ridged surface from high above. The largest feature is a circular ringed structure with high walled sides and a lower central peak: the entire surface out to the horizon is filled with similar structures that are smaller and overlapping.
	Lunar crater Daedalus on the Moon's far side
	
	A major geologic process that has affected the Moon's surface is impact cratering,[99] with craters formed when asteroids and comets collide with the lunar surface. There are estimated to be roughly 300,000 craters wider than 1 km (0.6 mi) on the Moon's near side.[100] The lunar geologic timescale is based on the most prominent impact events, including Nectaris, Imbrium, and Orientale; structures characterized by multiple rings of uplifted material, between hundreds and thousands of kilometers in diameter and associated with a broad apron of ejecta deposits that form a regional stratigraphic horizon.[101] The lack of an atmosphere, weather, and recent geological processes mean that many of these craters are well-preserved. Although only a few multi-ring basins have been definitively dated, they are useful for assigning relative ages. Because impact craters accumulate at a nearly constant rate, counting the number of craters per unit area can be used to estimate the age of the surface.[101] The radiometric ages of impact-melted rocks collected during the Apollo missions cluster between 3.8 and 4.1 billion years old: this has been used to propose a Late Heavy Bombardment period of increased impacts.[102]
	
	Blanketed on top of the Moon's crust is a highly comminuted (broken into ever smaller particles) and impact gardened surface layer called regolith, formed by impact processes. The finer regolith, the lunar soil of silicon dioxide glass, has a texture resembling snow and a scent resembling spent gunpowder.[103] The regolith of older surfaces is generally thicker than for younger surfaces: it varies in thickness from 10–20 km (6.2–12.4 mi) in the highlands and 3–5 km (1.9–3.1 mi) in the maria.[104] Beneath the finely comminuted regolith layer is the megaregolith, a layer of highly fractured bedrock many kilometers thick.[105]
	
	High-resolution images from the Lunar Reconnaissance Orbiter in the 2010s show a contemporary crater-production rate significantly higher than was previously estimated. A secondary cratering process caused by distal ejecta is thought to churn the top two centimeters of regolith on a timescale of 81,000 years.[106][107] This rate is 100 times faster than the rate computed from models based solely on direct micrometeorite impacts.[108]
	Lunar swirls at Reiner Gamma
	Gravitational field
	Main article: Gravity of the Moon
	GRAIL's gravity map of the Moon
	
	The gravitational field of the Moon has been measured through tracking the Doppler shift of radio signals emitted by orbiting spacecraft. The main lunar gravity features are mascons, large positive gravitational anomalies associated with some of the giant impact basins, partly caused by the dense mare basaltic lava flows that fill those basins.[109][110] The anomalies greatly influence the orbit of spacecraft about the Moon. There are some puzzles: lava flows by themselves cannot explain all of the gravitational signature, and some mascons exist that are not linked to mare volcanism.[111]
	Lunar swirls
	Main article: Lunar swirls
	
	Lunar swirls are enigmatic features found across the Moon's surface. They are characterized by a high albedo, appear optically immature (i.e. the optical characteristics of a relatively young regolith), and have often a sinuous shape. Their shape is often accentuated by low albedo regions that wind between the bright swirls. They are located in places with enhanced surface magnetic fields and many are located at the antipodal point of major impacts. Well known swirls include the Reiner Gamma feature and Mare Ingenii. They are hypothesized to be areas that have been partially shielded from the solar wind, resulting in slower space weathering.[112]
	Presence of water
	Main article: Lunar water
	
	Liquid water cannot persist on the lunar surface. When exposed to solar radiation, water quickly decomposes through a process known as photodissociation and is lost to space. However, since the 1960s, scientists have hypothesized that water ice may be deposited by impacting comets or possibly produced by the reaction of oxygen-rich lunar rocks, and hydrogen from solar wind, leaving traces of water which could possibly persist in cold, permanently shadowed craters at either pole on the Moon.[113][114] Computer simulations suggest that up to 14,000 km2 (5,400 sq mi) of the surface may be in permanent shadow.[115] The presence of usable quantities of water on the Moon is an important factor in rendering lunar habitation as a cost-effective plan; the alternative of transporting water from Earth would be prohibitively expensive.[116]
	
	In years since, signatures of water have been found to exist on the lunar surface.[117] In 1994, the bistatic radar experiment located on the Clementine spacecraft, indicated the existence of small, frozen pockets of water close to the surface. However, later radar observations by Arecibo, suggest these findings may rather be rocks ejected from young impact craters.[118] In 1998, the neutron spectrometer on the Lunar Prospector spacecraft showed that high concentrations of hydrogen are present in the first meter of depth in the regolith near the polar regions.[119] Volcanic lava beads, brought back to Earth aboard Apollo 15, showed small amounts of water in their interior.[120]
	
	The 2008 Chandrayaan-1 spacecraft has since confirmed the existence of surface water ice, using the on-board Moon Mineralogy Mapper. The spectrometer observed absorption lines common to hydroxyl, in reflected sunlight, providing evidence of large quantities of water ice, on the lunar surface. The spacecraft showed that concentrations may possibly be as high as 1,000 ppm.[121] Using the mapper's reflectance spectra, indirect lighting of areas in shadow confirmed water ice within 20° latitude of both poles in 2018.[122] In 2009, LCROSS sent a 2,300 kg (5,100 lb) impactor into a permanently shadowed polar crater, and detected at least 100 kg (220 lb) of water in a plume of ejected material.[123][124] Another examination of the LCROSS data showed the amount of detected water to be closer to 155 ± 12 kg (342 ± 26 lb).[125]
	
	In May 2011, 615–1410 ppm water in melt inclusions in lunar sample 74220 was reported,[126] the famous high-titanium "orange glass soil" of volcanic origin collected during the Apollo 17 mission in 1972. The inclusions were formed during explosive eruptions on the Moon approximately 3.7 billion years ago. This concentration is comparable with that of magma in Earth's upper mantle. Although of considerable selenological interest, this announcement affords little comfort to would-be lunar colonists – the sample originated many kilometers below the surface, and the inclusions are so difficult to access that it took 39 years to find them with a state-of-the-art ion microprobe instrument.
	
	Analysis of the findings of the Moon Mineralogy Mapper (M3) revealed in August 2018 for the first time "definitive evidence" for water-ice on the lunar surface.[127][128] The data revealed the distinct reflective signatures of water-ice, as opposed to dust and other reflective substances.[129] The ice deposits were found on the North and South poles, although it is more abundant in the South, where water is trapped in permanently shadowed craters and crevices, allowing it to persist as ice on the surface since they are shielded from the sun.[127][129]
	
	In October 2020, astronomers reported detecting molecular water on the sunlit surface of the Moon by several independent spacecraft, including the Stratospheric Observatory for Infrared Astronomy (SOFIA).[130][131][132][133]
	Surface conditions
	
	The surface of the Moon is an extreme environment with temperatures that range from 140 °C down to −171 °C, an atmospheric pressure of 10−10 Pa, and high levels of ionizing radiation from the Sun and cosmic rays. The exposed surfaces of spacecraft are considered unlikely to harbor bacterial spores after just one lunar orbit.[134] The surface gravity of the Moon is approximately 1.625 m/s2, about 16.6% that on Earth's surface or 0.166 ɡ.[4]
	Atmosphere
	Main article: Atmosphere of the Moon
	Sketch by the Apollo 17 astronauts. The lunar atmosphere was later studied by LADEE.[135][136]
	
	The Moon has an atmosphere so tenuous as to be nearly vacuum, with a total mass of less than 10 tonnes (9.8 long tons; 11 short tons).[137] The surface pressure of this small mass is around 3 × 10−15 atm (0.3 nPa); it varies with the lunar day. Its sources include outgassing and sputtering, a product of the bombardment of lunar soil by solar wind ions.[14][138] Elements that have been detected include sodium and potassium, produced by sputtering (also found in the atmospheres of Mercury and Io); helium-4 and neon[139] from the solar wind; and argon-40, radon-222, and polonium-210, outgassed after their creation by radioactive decay within the crust and mantle.[140][141] The absence of such neutral species (atoms or molecules) as oxygen, nitrogen, carbon, hydrogen and magnesium, which are present in the regolith, is not understood.[140] Water vapor has been detected by Chandrayaan-1 and found to vary with latitude, with a maximum at ~60–70 degrees; it is possibly generated from the sublimation of water ice in the regolith.[142] These gases either return into the regolith because of the Moon's gravity or are lost to space, either through solar radiation pressure or, if they are ionized, by being swept away by the solar wind's magnetic field.[140]
	
	Studies of Moon magma samples retrieved by the Apollo missions demonstrate that the Moon had once possessed a relatively thick atmosphere for a period of 70 million years between 3 and 4 billion years ago. This atmosphere, sourced from gases ejected from lunar volcanic eruptions, was twice the thickness of that of present-day Mars. The ancient lunar atmosphere was eventually stripped away by solar winds and dissipated into space.[143]
	Dust
	
	A permanent Moon dust cloud exists around the Moon, generated by small particles from comets. Estimates are 5 tons of comet particles strike the Moon's surface every 24 hours, resulting in the ejection of dust particles. The dust stays above the Moon approximately 10 minutes, taking 5 minutes to rise, and 5 minutes to fall. On average, 120 kilograms of dust are present above the Moon, rising up to 100 kilometers above the surface. Dust counts made by LADEE's Lunar Dust EXperiment (LDEX) found particle counts peaked during the Geminid, Quadrantid, Northern Taurid, and Omicron Centaurid meteor showers, when the Earth, and Moon pass through comet debris. The lunar dust cloud is asymmetric, being more dense near the boundary between the Moon's dayside and nightside.[144][145]
	Earth–Moon system
	
	See also: Satellite system (astronomy), Other moons of Earth, and Double planet
	Lunar distance
	Main article: Lunar distance (astronomy)
	These paragraphs are an excerpt from Lunar distance (astronomy).[edit]
	The instantaneous Earth–Moon distance, or distance to the Moon, is the distance from the center of Earth to the center of the Moon. Lunar distance (LD or Δ ⊕ L {\textstyle \Delta _{\oplus L}} {\textstyle \Delta _{\oplus L}}), or Earth–Moon characteristic distance, is a unit of measure in astronomy. More technically, it is the semi-major axis of the geocentric lunar orbit. The lunar distance is approximately 400,000 km, which is a quarter of a million miles or 1.28 light-seconds. This is roughly thirty times Earth's diameter.
	
		Scale model of the Earth–Moon system: Sizes and distances are to scale.
	
	Minimum, mean and maximum distances of the Moon from Earth with its angular diameter as seen from Earth's surface, to scale
	Orbit
	Main articles: Orbit of the Moon and Lunar theory
	Earth has a pronounced axial tilt; the Moon's orbit is not perpendicular to Earth's axis, but lies close to Earth's orbital plane.
	Earth–Moon system (schematic)
	
	Because of tidal locking, the rotation of the Moon around its own axis is synchronous to its orbital period around the Earth. The Moon makes a complete orbit around Earth with respect to the fixed stars about once every 27.3 days,[g] its sidereal period. However, because Earth is moving in its orbit around the Sun at the same time, it takes slightly longer for the Moon to show the same phase to Earth, which is about 29.5 days;[h] its synodic period.[80][146]
	
	Unlike most satellites of other planets, the Moon orbits closer to the ecliptic plane than to the planet's equatorial plane. The Moon's orbit is subtly perturbed by the Sun and Earth in many small, complex and interacting ways. For example, the plane of the Moon's orbit gradually rotates once every 18.61 years,[147] which affects other aspects of lunar motion. These follow-on effects are mathematically described by Cassini's laws.[148]
	
	The Moon's axial tilt with respect to the ecliptic is only 1.5427°,[8][149] much less than the 23.44° of Earth. Because of this, the Moon's solar illumination varies much less with season, and topographical details play a crucial role in seasonal effects.[150] From images taken by Clementine in 1994, it appears that four mountainous regions on the rim of the crater Peary at the Moon's north pole may remain illuminated for the entire lunar day, creating peaks of eternal light. No such regions exist at the south pole. Similarly, there are places that remain in permanent shadow at the bottoms of many polar craters,[115] and these "craters of eternal darkness" are extremely cold: Lunar Reconnaissance Orbiter measured the lowest summer temperatures in craters at the southern pole at 35 K (−238 °C; −397 °F)[151] and just 26 K (−247 °C; −413 °F) close to the winter solstice in the north polar crater Hermite. This is the coldest temperature in the Solar System ever measured by a spacecraft, colder even than the surface of Pluto.[150] Average temperatures of the Moon's surface are reported, but temperatures of different areas will vary greatly depending upon whether they are in sunlight or shadow.[152]
	Relative size
	DSCOVR satellite sees the Moon passing in front of Earth
	
	The Moon is an exceptionally large natural satellite relative to Earth: Its diameter is more than a quarter and its mass is 1/81 of Earth's.[80] It is the largest moon in the Solar System relative to the size of its planet,[i] though Charon is larger relative to the dwarf planet Pluto, at 1/9 Pluto's mass.[j][153] The Earth and the Moon's barycentre, their common center of mass, is located 1,700 km (1,100 mi) (about a quarter of Earth's radius) beneath the Earth's surface.
	
	The Earth revolves around the Earth-Moon barycentre once a sidereal month, with 1/81 the speed of the Moon, or about 12.5 metres (41 ft) per second. This motion is superimposed on the much larger revolution of the Earth around the Sun at a speed of about 30 kilometres (19 mi) per second.
	
	The surface area of the Moon is slightly less than the areas of North and South America combined.
	Appearance from Earth
	See also: Lunar observation, Lunar phase, Moonlight, and Earthlight (astronomy)
	
	The synchronous rotation of the Moon as it orbits the Earth results in it always keeping nearly the same face turned towards the planet. However, because of the effect of libration, about 59% of the Moon's surface can actually be seen from Earth. The side of the Moon that faces Earth is called the near side, and the opposite the far side. The far side is often inaccurately called the "dark side", but it is in fact illuminated as often as the near side: once every 29.5 Earth days. During new moon, the near side is dark.[154]
	
	The Moon originally rotated at a faster rate, but early in its history its rotation slowed and became tidally locked in this orientation as a result of frictional effects associated with tidal deformations caused by Earth.[155] With time, the energy of rotation of the Moon on its axis was dissipated as heat, until there was no rotation of the Moon relative to Earth. In 2016, planetary scientists using data collected on the 1998-99 NASA Lunar Prospector mission, found two hydrogen-rich areas (most likely former water ice) on opposite sides of the Moon. It is speculated that these patches were the poles of the Moon billions of years ago before it was tidally locked to Earth.[156]
	During the lunar phases, only portions of the Moon can be observed from Earth.
	
	The Moon has an exceptionally low albedo, giving it a reflectance that is slightly brighter than that of worn asphalt. Despite this, it is the brightest object in the sky after the Sun.[80][k] This is due partly to the brightness enhancement of the opposition surge; the Moon at quarter phase is only one-tenth as bright, rather than half as bright, as at full moon.[157] Additionally, color constancy in the visual system recalibrates the relations between the colors of an object and its surroundings, and because the surrounding sky is comparatively dark, the sunlit Moon is perceived as a bright object. The edges of the full moon seem as bright as the center, without limb darkening, because of the reflective properties of lunar soil, which retroreflects light more towards the Sun than in other directions. The Moon does appear larger when close to the horizon, but this is a purely psychological effect, known as the Moon illusion, first described in the 7th century BC.[158] The full Moon's angular diameter is about 0.52° (on average) in the sky, roughly the same apparent size as the Sun (see § Eclipses).
	
	The Moon's highest altitude at culmination varies by its phase and time of year. The full moon is highest in the sky during winter (for each hemisphere). The orientation of the Moon's crescent also depends on the latitude of the viewing location; an observer in the tropics can see a smile-shaped crescent Moon.[159] The Moon is visible for two weeks every 27.3 days at the North and South Poles. Zooplankton in the Arctic use moonlight when the Sun is below the horizon for months on end.[160]
	A full moon appears as a half moon during an eclipse moonset over the High Desert in California, on the morning of the Trifecta: Full moon, Supermoon, Lunar eclipse, January 2018 lunar eclipse
	
	The distance between the Moon and Earth varies from around 356,400 km (221,500 mi) to 406,700 km (252,700 mi) at perigee (closest) and apogee (farthest), respectively. On 14 November 2016, it was closer to Earth when at full phase than it has been since 1948, 14% closer than its farthest position in apogee.[161] Reported as a "supermoon", this closest point coincided within an hour of a full moon, and it was 30% more luminous than when at its greatest distance because its angular diameter is 14% greater and 1.14 2 ≈ 1.30 {\displaystyle \scriptstyle 1.14^{2}\approx 1.30} \scriptstyle 1.14^{2}\approx 1.30.[162][163][164] At lower levels, the human perception of reduced brightness as a percentage is provided by the following formula:[165][166]
	
		perceived reduction % = 100 × actual reduction % 100 {\displaystyle {\text{perceived reduction}}\%=100\times {\sqrt {{\text{actual reduction}}\% \over 100}}} {\displaystyle {\text{perceived reduction}}\%=100\times {\sqrt {{\text{actual reduction}}\% \over 100}}}
	
	When the actual reduction is 1.00 / 1.30, or about 0.770, the perceived reduction is about 0.877, or 1.00 / 1.14. This gives a maximum perceived increase of 14% between apogee and perigee moons of the same phase.[167]
	
	There has been historical controversy over whether features on the Moon's surface change over time. Today, many of these claims are thought to be illusory, resulting from observation under different lighting conditions, poor astronomical seeing, or inadequate drawings. However, outgassing does occasionally occur and could be responsible for a minor percentage of the reported lunar transient phenomena. Recently, it has been suggested that a roughly 3 km (1.9 mi) diameter region of the lunar surface was modified by a gas release event about a million years ago.[168][169]
	
	The Moon's appearance, like the Sun's, can be affected by Earth's atmosphere. Common optical effects are the 22° halo ring, formed when the Moon's light is refracted through the ice crystals of high cirrostratus clouds, and smaller coronal rings when the Moon is seen through thin clouds.[170]
	The monthly changes in the angle between the direction of sunlight and view from Earth, and the phases of the Moon that result, as viewed from the Northern Hemisphere. The Earth–Moon distance is not to scale.
	
	The illuminated area of the visible sphere (degree of illumination) is given by ( 1 − cos ⁡ e ) / 2 = sin 2 ⁡ ( e / 2 ) {\displaystyle (1-\cos e)/2=\sin ^{2}(e/2)} {\displaystyle (1-\cos e)/2=\sin ^{2}(e/2)}, where e {\displaystyle e} e is the elongation (i.e., the angle between Moon, the observer on Earth, and the Sun).
	Eclipses
	Main articles: Solar eclipse, Lunar eclipse, and Eclipse cycle
	The Moon, tinted reddish, during a lunar eclipse
	The fiercely bright disk of the Sun is completely obscured by the exact fit of the disk of the dark, non-illuminated Moon, leaving only the radial, fuzzy, glowing coronal filaments of the Sun around the edge.
	The bright disk of the Sun, showing many coronal filaments, flares and grainy patches in the wavelength of this image, is partly obscured by a small dark disk: here, the Moon covers less than a fifteenth of the Sun.
	From Earth, the Moon and the Sun appear the same size, as seen in the 1999 solar eclipse (left), whereas from the STEREO-B spacecraft in an Earth-trailing orbit, the Moon appears much smaller than the Sun (right).[171]
	
	Eclipses only occur when the Sun, Earth, and Moon are all in a straight line (termed "syzygy"). Solar eclipses occur at new moon, when the Moon is between the Sun and Earth. In contrast, lunar eclipses occur at full moon, when Earth is between the Sun and Moon. The apparent size of the Moon is roughly the same as that of the Sun, with both being viewed at close to one-half a degree wide. The Sun is much larger than the Moon but it is the vastly greater distance that gives it the same apparent size as the much closer and much smaller Moon from the perspective of Earth. The variations in apparent size, due to the non-circular orbits, are nearly the same as well, though occurring in different cycles. This makes possible both total (with the Moon appearing larger than the Sun) and annular (with the Moon appearing smaller than the Sun) solar eclipses.[172] In a total eclipse, the Moon completely covers the disc of the Sun and the solar corona becomes visible to the naked eye. Because the distance between the Moon and Earth is very slowly increasing over time,[173] the angular diameter of the Moon is decreasing. Also, as it evolves toward becoming a red giant, the size of the Sun, and its apparent diameter in the sky, are slowly increasing.[l] The combination of these two changes means that hundreds of millions of years ago, the Moon would always completely cover the Sun on solar eclipses, and no annular eclipses were possible. Likewise, hundreds of millions of years in the future, the Moon will no longer cover the Sun completely, and total solar eclipses will not occur.[174]
	
	Because the Moon's orbit around Earth is inclined by about 5.145° (5° 9') to the orbit of Earth around the Sun, eclipses do not occur at every full and new moon. For an eclipse to occur, the Moon must be near the intersection of the two orbital planes.[175] The periodicity and recurrence of eclipses of the Sun by the Moon, and of the Moon by Earth, is described by the saros, which has a period of approximately 18 years.[176]
	
	Because the Moon continuously blocks the view of a half-degree-wide circular area of the sky,[m][177] the related phenomenon of occultation occurs when a bright star or planet passes behind the Moon and is occulted: hidden from view. In this way, a solar eclipse is an occultation of the Sun. Because the Moon is comparatively close to Earth, occultations of individual stars are not visible everywhere on the planet, nor at the same time. Because of the precession of the lunar orbit, each year different stars are occulted.[178]
	Tidal effects
	Main articles: Tidal force, Tidal acceleration, Tide, and Theory of tides
	Over one lunar month more than half of the Moon's surface can be seen from Earth's surface.
	The libration of the Moon over a single lunar month. Also visible is the slight variation in the Moon's visual size from Earth.
	
	The gravitational attraction that masses have for one another decreases inversely with the square of the distance of those masses from each other. As a result, the slightly greater attraction that the Moon has for the side of Earth closest to the Moon, as compared to the part of the Earth opposite the Moon, results in tidal forces. Tidal forces affect both the Earth's crust and oceans.
	
	The most obvious effect of tidal forces is to cause two bulges in the Earth's oceans, one on the side facing the Moon and the other on the side opposite. This results in elevated sea levels called ocean tides.[173] As the Earth rotates on its axis, one of the ocean bulges (high tide) is held in place "under" the Moon, while another such tide is opposite. As a result, there are two high tides, and two low tides in about 24 hours.[173] Since the Moon is orbiting the Earth in the same direction of the Earth's rotation, the high tides occur about every 12 hours and 25 minutes; the 25 minutes is due to the Moon's time to orbit the Earth. The Sun has the same tidal effect on the Earth, but its forces of attraction are only 40% that of the Moon's; the Sun's and Moon's interplay is responsible for spring and neap tides.[173] If the Earth were a water world (one with no continents) it would produce a tide of only one meter, and that tide would be very predictable, but the ocean tides are greatly modified by other effects: the frictional coupling of water to Earth's rotation through the ocean floors, the inertia of water's movement, ocean basins that grow shallower near land, the sloshing of water between different ocean basins.[179] As a result, the timing of the tides at most points on the Earth is a product of observations that are explained, incidentally, by theory.
	
	While gravitation causes acceleration and movement of the Earth's fluid oceans, gravitational coupling between the Moon and Earth's solid body is mostly elastic and plastic. The result is a further tidal effect of the Moon on the Earth that causes a bulge of the solid portion of the Earth nearest the Moon. Delays in the tidal peaks of both ocean and solid-body tides cause torque in opposition to the Earth's rotation. This "drains" angular momentum and rotational kinetic energy from Earth's rotation, slowing the Earth's rotation.[173][180] That angular momentum, lost from the Earth, is transferred to the Moon in a process (confusingly known as tidal acceleration), which lifts the Moon into a higher orbit and results in its lower orbital speed about the Earth. Thus the distance between Earth and Moon is increasing, and the Earth's rotation is slowing in reaction.[180] Measurements from laser reflectors left during the Apollo missions (lunar ranging experiments) have found that the Moon's distance increases by 38 mm (1.5 in) per year (roughly the rate at which human fingernails grow).[181][182][183] Atomic clocks also show that Earth's day lengthens by about 17 microseconds every year,[184][185][186] slowly increasing the rate at which UTC is adjusted by leap seconds. This tidal drag would continue until the rotation of Earth and the orbital period of the Moon matched, creating mutual tidal locking between the two and suspending the Moon over one meridian (this is currently the case with Pluto and its moon Charon). However, the Sun will become a red giant engulfing the Earth-Moon system long before this occurrence.[187][188]
	
	In a like manner, the lunar surface experiences tides of around 10 cm (4 in) amplitude over 27 days, with three components: a fixed one due to Earth, because they are in synchronous rotation, a variable tide due to orbital eccentricity and inclination, and a small varying component from the Sun.[180] The Earth-induced variable component arises from changing distance and libration, a result of the Moon's orbital eccentricity and inclination (if the Moon's orbit were perfectly circular and un-inclined, there would only be solar tides).[180] Libration also changes the angle from which the Moon is seen, allowing a total of about 59% of its surface to be seen from Earth over time.[80] The cumulative effects of stress built up by these tidal forces produces moonquakes. Moonquakes are much less common and weaker than are earthquakes, although moonquakes can last for up to an hour – significantly longer than terrestrial quakes – because of scattering of the seismic vibrations in the dry fragmented upper crust. The existence of moonquakes was an unexpected discovery from seismometers placed on the Moon by Apollo astronauts from 1969 through 1972.[189]
	
	According to recent research, scientists suggest that the Moon's influence on the Earth may contribute to maintaining Earth's magnetic field.[190]
	Observation and exploration
	Main articles: Exploration of the Moon, List of spacecraft that orbited the Moon, List of missions to the Moon, and List of lunar probes
	See also: Timeline of Solar System exploration
	Before spaceflight
	Main article: Exploration of the Moon: Before spaceflight
	On an open folio page is a carefully drawn disk of the full moon. In the upper corners of the page are waving banners held aloft by pairs of winged cherubs. In the lower left page corner a cherub assists another to measure distances with a pair of compasses; in the lower right corner a cherub views the main map through a handheld telescope, whereas another, kneeling, peers at the map from over a low cloth-draped table.
	Map of the Moon by Johannes Hevelius from his Selenographia (1647), the first map to include the libration zones
	
	One of the earliest-discovered possible depictions of the Moon is a 5000-year-old rock carving Orthostat 47 at Knowth, Ireland.[191][192]
	
	Understanding of the Moon's cycles was an early development of astronomy: by the 5th century BC, Babylonian astronomers had recorded the 18-year Saros cycle of lunar eclipses,[193] and Indian astronomers had described the Moon's monthly elongation.[194] The Chinese astronomer Shi Shen (fl. 4th century BC) gave instructions for predicting solar and lunar eclipses.[195]: 411  Later, the physical form of the Moon and the cause of moonlight became understood. The ancient Greek philosopher Anaxagoras (d. 428 BC) reasoned that the Sun and Moon were both giant spherical rocks, and that the latter reflected the light of the former.[196][195]: 227  Although the Chinese of the Han Dynasty believed the Moon to be energy equated to qi, their 'radiating influence' theory also recognized that the light of the Moon was merely a reflection of the Sun, and Jing Fang (78–37 BC) noted the sphericity of the Moon.[195]: 413–414  In the 2nd century AD, Lucian wrote the novel A True Story, in which the heroes travel to the Moon and meet its inhabitants. In 499 AD, the Indian astronomer Aryabhata mentioned in his Aryabhatiya that reflected sunlight is the cause of the shining of the Moon.[197] The astronomer and physicist Alhazen (965–1039) found that sunlight was not reflected from the Moon like a mirror, but that light was emitted from every part of the Moon's sunlit surface in all directions.[198] Shen Kuo (1031–1095) of the Song dynasty created an allegory equating the waxing and waning of the Moon to a round ball of reflective silver that, when doused with white powder and viewed from the side, would appear to be a crescent.[195]: 415–416 
	Galileo's sketches of the Moon from Sidereus Nuncius
	
	In Aristotle's (384–322 BC) description of the universe, the Moon marked the boundary between the spheres of the mutable elements (earth, water, air and fire), and the imperishable stars of aether, an influential philosophy that would dominate for centuries.[199] However, in the 2nd century BC, Seleucus of Seleucia correctly theorized that tides were due to the attraction of the Moon, and that their height depends on the Moon's position relative to the Sun.[200] In the same century, Aristarchus computed the size and distance of the Moon from Earth, obtaining a value of about twenty times the radius of Earth for the distance. These figures were greatly improved by Ptolemy (90–168 AD): his values of a mean distance of 59 times Earth's radius and a diameter of 0.292 Earth diameters were close to the correct values of about 60 and 0.273 respectively.[201] Archimedes (287–212 BC) designed a planetarium that could calculate the motions of the Moon and other objects in the Solar System.[202]
	
	During the Middle Ages, before the invention of the telescope, the Moon was increasingly recognised as a sphere, though many believed that it was "perfectly smooth".[203]
	
	In 1609, Galileo Galilei used an early telescope to make drawings of the Moon for his book Sidereus Nuncius, and deduced that it was not smooth but had mountains and craters. Thomas Harriot had made, but not published such drawings a few months earlier. Telescopic mapping of the Moon followed: later in the 17th century, the efforts of Giovanni Battista Riccioli and Francesco Maria Grimaldi led to the system of naming of lunar features in use today. The more exact 1834–1836 Mappa Selenographica of Wilhelm Beer and Johann Heinrich Mädler, and their associated 1837 book Der Mond, the first trigonometrically accurate study of lunar features, included the heights of more than a thousand mountains, and introduced the study of the Moon at accuracies possible in earthly geography.[204] Lunar craters, first noted by Galileo, were thought to be volcanic until the 1870s proposal of Richard Proctor that they were formed by collisions.[80] This view gained support in 1892 from the experimentation of geologist Grove Karl Gilbert, and from comparative studies from 1920 to the 1940s,[205] leading to the development of lunar stratigraphy, which by the 1950s was becoming a new and growing branch of astrogeology.[80]
	1959–1970s
	See also: Space Race and Moon landing
	
	Between the first human arrival with the robotic Soviet Luna program in 1958, to the 1970s with the last Missions of the crewed U.S. Apollo landings and last Luna mission in 1976, the Cold War-inspired Space Race between the Soviet Union and the U.S. led to an acceleration of interest in exploration of the Moon. Once launchers had the necessary capabilities, these nations sent uncrewed probes on both flyby and impact/lander missions.
	Soviet missions
	Main articles: Luna program and Lunokhod programme
	First view in history of the far side of the Moon, taken by Luna 3, 7 October 1959
	
	Spacecraft from the Soviet Union's Luna program were the first to accomplish a number of goals: following three unnamed, failed missions in 1958,[206] the first human-made object to escape Earth's gravity and pass near the Moon was Luna 1; the first human-made object to impact the lunar surface was Luna 2, and the first photographs of the normally occluded far side of the Moon were made by Luna 3, all in 1959. The first spacecraft to perform a successful lunar soft landing was Luna 9 and the first vehicle to orbit the Moon was Luna 10, both in 1966.[80] Rock and soil samples were brought back to Earth by three Luna sample return missions (Luna 16 in 1970, Luna 20 in 1972, and Luna 24 in 1976), which returned 0.3 kg total.[207] Luna 17 deployed the first lunar rover, Lunokhod 1, in 1970.
	United States missions
	Main articles: Apollo program and Moon landing
	The small blue-white semicircle of Earth, almost glowing with color in the blackness of space, rising over the limb of the desolate, cratered surface of the Moon.
	Earthrise (Apollo 8, 1968, taken by William Anders)
	Moon rock (Lunar basalt 70017, Apollo 17, 1972)
	
	During the late 1950s at the height of the Cold War, the United States Army conducted a classified feasibility study that proposed the construction of a staffed military outpost on the Moon called Project Horizon with the potential to conduct a wide range of missions from scientific research to nuclear Earth bombardment. The study included the possibility of conducting a lunar-based nuclear test.[208][209] The Air Force, which at the time was in competition with the Army for a leading role in the space program, developed its own similar plan called Lunex.[210][211][208] However, both these proposals were ultimately passed over as the space program was largely transferred from the military to the civilian agency NASA.[211]
	
	Following President John F. Kennedy's 1961 commitment to a manned Moon landing before the end of the decade, the United States, under NASA leadership, launched a series of uncrewed probes to develop an understanding of the lunar surface in preparation for human missions: the Jet Propulsion Laboratory's Ranger program produced the first close-up pictures; the Lunar Orbiter program produced maps of the entire Moon; the Surveyor program landed its first spacecraft four months after Luna 9. The crewed Apollo program was developed in parallel; after a series of uncrewed and crewed tests of the Apollo spacecraft in Earth orbit, and spurred on by a potential Soviet lunar human landing, in 1968 Apollo 8 made the first human mission to lunar orbit. The subsequent landing of the first humans on the Moon in 1969 is seen by many as the culmination of the Space Race.[212]
	Neil Armstrong working at the Lunar Module Eagle during Apollo 11 (1969)
		
	"That's one small step ..."
	Menu
	0:00
	Problems playing this file? See media help.
	
	Neil Armstrong became the first person to walk on the Moon as the commander of the American mission Apollo 11 by first setting foot on the Moon at 02:56 UTC on 21 July 1969.[213] An estimated 500 million people worldwide watched the transmission by the Apollo TV camera, the largest television audience for a live broadcast at that time.[214][215] The Apollo missions 11 to 17 (except Apollo 13, which aborted its planned lunar landing) removed 380.05 kilograms (837.87 lb) of lunar rock and soil in 2,196 separate samples.[216] The American Moon landing and return was enabled by considerable technological advances in the early 1960s, in domains such as ablation chemistry, software engineering, and atmospheric re-entry technology, and by highly competent management of the enormous technical undertaking.[217][218]
	
	Scientific instrument packages were installed on the lunar surface during all the Apollo landings. Long-lived instrument stations, including heat flow probes, seismometers, and magnetometers, were installed at the Apollo 12, 14, 15, 16, and 17 landing sites. Direct transmission of data to Earth concluded in late 1977 because of budgetary considerations,[219][220] but as the stations' lunar laser ranging corner-cube retroreflector arrays are passive instruments, they are still being used. Ranging to the stations is routinely performed from Earth-based stations with an accuracy of a few centimeters, and data from this experiment are being used to place constraints on the size of the lunar core.[221]
	1970s – present
	
	In the 1970s, after the Moon race, the focus of astronautic exploration shifted, as probes like Pioneer 10 and the Voyager program were sent towards the outer solar system. Years of near lunar quietude followed, only broken by a beginning internationalization of space and the Moon through, for example, the negotiation of the Moon treaty.
	
	Since the 1990s, many more countries have become involved in direct exploration of the Moon. In 1990, Japan became the third country to place a spacecraft into lunar orbit with its Hiten spacecraft. The spacecraft released a smaller probe, Hagoromo, in lunar orbit, but the transmitter failed, preventing further scientific use of the mission.[222] In 1994, the U.S. sent the joint Defense Department/NASA spacecraft Clementine to lunar orbit. This mission obtained the first near-global topographic map of the Moon, and the first global multispectral images of the lunar surface.[223] This was followed in 1998 by the Lunar Prospector mission, whose instruments indicated the presence of excess hydrogen at the lunar poles, which is likely to have been caused by the presence of water ice in the upper few meters of the regolith within permanently shadowed craters.[224]
	As viewed by Chandrayaan-1's NASA Moon Mineralogy Mapper equipment, on the right, the first time discovered water-rich minerals (light blue), shown around a small crater from which it was ejected.
	
	The European spacecraft SMART-1, the second ion-propelled spacecraft, was in lunar orbit from 15 November 2004 until its lunar impact on 3 September 2006, and made the first detailed survey of chemical elements on the lunar surface.[225]
	
	The ambitious Chinese Lunar Exploration Program began with Chang'e 1, which successfully orbited the Moon from 5 November 2007 until its controlled lunar impact on 1 March 2009.[226] It obtained a full image map of the Moon. Chang'e 2, beginning in October 2010, reached the Moon more quickly, mapped the Moon at a higher resolution over an eight-month period, then left lunar orbit for an extended stay at the Earth–Sun L2 Lagrangian point, before finally performing a flyby of asteroid 4179 Toutatis on 13 December 2012, and then heading off into deep space. On 14 December 2013, Chang'e 3 landed a lunar lander onto the Moon's surface, which in turn deployed a lunar rover, named Yutu (Chinese: 玉兔; literally "Jade Rabbit"). This was the first lunar soft landing since Luna 24 in 1976, and the first lunar rover mission since Lunokhod 2 in 1973. Another rover mission (Chang'e 4) was launched in 2019, becoming the first ever spacecraft to land on the Moon's far side. China intends to following this up with a sample return mission (Chang'e 5) in 2020.[227]
	
	Between 4 October 2007 and 10 June 2009, the Japan Aerospace Exploration Agency's Kaguya (Selene) mission, a lunar orbiter fitted with a high-definition video camera, and two small radio-transmitter satellites, obtained lunar geophysics data and took the first high-definition movies from beyond Earth orbit.[228][229] India's first lunar mission, Chandrayaan-1, orbited from 8 November 2008 until loss of contact on 27 August 2009, creating a high-resolution chemical, mineralogical and photo-geological map of the lunar surface, and confirming the presence of water molecules in lunar soil.[230] The Indian Space Research Organisation planned to launch Chandrayaan-2 in 2013, which would have included a Russian robotic lunar rover.[231][232] However, the failure of Russia's Fobos-Grunt mission has delayed this project, and was launched on 22 July 2019. The lander Vikram attempted to land on the lunar south pole region on 6 September, but lost the signal in 2.1 km (1.3 mi). What happened after that is unknown.
	
	The U.S. co-launched the Lunar Reconnaissance Orbiter (LRO) and the LCROSS impactor and follow-up observation orbiter on 18 June 2009; LCROSS completed its mission by making a planned and widely observed impact in the crater Cabeus on 9 October 2009,[233] whereas LRO is currently in operation, obtaining precise lunar altimetry and high-resolution imagery. In November 2011, the LRO passed over the large and bright crater Aristarchus. NASA released photos of the crater on 25 December 2011.[234]
	
	Two NASA GRAIL spacecraft began orbiting the Moon around 1 January 2012,[235] on a mission to learn more about the Moon's internal structure. NASA's LADEE probe, designed to study the lunar exosphere, achieved orbit on 6 October 2013.
	Future
	See also: List of proposed missions to the Moon
	
	Upcoming lunar missions include Russia's Luna-Glob: an uncrewed lander with a set of seismometers, and an orbiter based on its failed Martian Fobos-Grunt mission.[236] Privately funded lunar exploration has been promoted by the Google Lunar X Prize, announced 13 September 2007, which offers US$20 million to anyone who can land a robotic rover on the Moon and meet other specified criteria.[237]
	
	NASA began to plan to resume human missions following the call by U.S. President George W. Bush on 14 January 2004 for a human mission to the Moon by 2019 and the construction of a lunar base by 2024.[238] The Constellation program was funded and construction and testing begun on a crewed spacecraft and launch vehicle,[239] and design studies for a lunar base.[240] That program was cancelled in 2010, however, and was eventually replaced with the Donald Trump supported Artemis program, which plans to return humans to the Moon by 2025.[241] India had also expressed its hope to send people to the Moon by 2020.[242]
	
	On 28 February 2018, SpaceX, Vodafone, Nokia and Audi announced a collaboration to install a 4G wireless communication network on the Moon, with the aim of streaming live footage on the surface to Earth.[243]
	
	Recent reports also indicate NASA's intent to send a woman astronaut to the Moon in their planned mid-2020s mission.[244]
	Planned commercial missions
	
	In 2007, the X Prize Foundation together with Google launched the Google Lunar X Prize to encourage commercial endeavors to the Moon. A prize of $20 million was to be awarded to the first private venture to get to the Moon with a robotic lander by the end of March 2018, with additional prizes worth $10 million for further milestones.[245][246] As of August 2016, 16 teams were reportedly participating in the competition.[247] In January 2018 the foundation announced that the prize would go unclaimed as none of the finalist teams would be able to make a launch attempt by the deadline.[248]
	
	In August 2016, the US government granted permission to US-based start-up Moon Express to land on the Moon.[249] This marked the first time that a private enterprise was given the right to do so. The decision is regarded as a precedent helping to define regulatory standards for deep-space commercial activity in the future. Previously, private companies were restricted to operating on or around Earth.[249]
	
	On 29 November 2018 NASA announced that nine commercial companies would compete to win a contract to send small payloads to the Moon in what is known as Commercial Lunar Payload Services. According to NASA administrator Jim Bridenstine, "We are building a domestic American capability to get back and forth to the surface of the moon.".[250]
	Human presence
	See also: Human presence in space
	Human impact
	See also: List of artificial objects on the Moon, Space art § Art in space, and Planetary protection § Category V
	Remains of human activity, Apollo 17's Lunar Surface Experiments Package
	
	Beside the traces of human activity on the Moon, there have been some intended permanent installations like the Moon Museum art piece, Apollo 11 goodwill messages, six Lunar plaques, the Fallen Astronaut memorial, and other artifacts.
	Infrastructure
	Main article: Moonbase
	See also: Space infrastructure, Tourism on the Moon, Lunar resources § Mining, and Colonization of the Moon
	A photo of the still in use reflector of the Lunar Laser Ranging Experiment of Apollo 11.
	
	Longterm missions continuing to be active are some orbiters such as the 2009-launched Lunar Reconnaissance Orbiter surveilling the Moon for future missions, as well as some Landers such as the 2013-launched Chang'e 3 with its Lunar Ultraviolet Telescope still operational.[251]
	
	There are several missions by different agencies and companies planned to establish a longterm human presence on the Moon, with the Lunar Gateway as the currently most advanced project as part of the Artemis program.
	Astronomy from the Moon
	Further information: Extraterrestrial sky § The Moon
	
	For many years, the Moon has been recognized as an excellent site for telescopes.[252] It is relatively nearby; astronomical seeing is not a concern; certain craters near the poles are permanently dark and cold, and thus especially useful for infrared telescopes; and radio telescopes on the far side would be shielded from the radio chatter of Earth.[253] The lunar soil, although it poses a problem for any moving parts of telescopes, can be mixed with carbon nanotubes and epoxies and employed in the construction of mirrors up to 50 meters in diameter.[254] A lunar zenith telescope can be made cheaply with an ionic liquid.[255]
	
	In April 1972, the Apollo 16 mission recorded various astronomical photos and spectra in ultraviolet with the Far Ultraviolet Camera/Spectrograph.[256]
	Living on the Moon
	
	Humans have stayed for days on the Moon, such as during Apollo 17.[257] One particular challenge for astronauts' daily life during their stay on the surface is the lunar dust sticking to their suits and being carried into their quarters. Subsequently, the dust was tasted and smelled by the astronauts, calling it the "Apollo aroma".[258] This contamination poses a danger since the fine lunar dust can cause health issues.[258]
	
	In 2019 at least one plant seed sprouted in an experiment, carried along with other small life from Earth on the Chang'e 4 lander in its Lunar Micro Ecosystem.[259]
	Legal status
	Main article: Space law
	
	Although Luna landers scattered pennants of the Soviet Union on the Moon, and U.S. flags were symbolically planted at their landing sites by the Apollo astronauts, no nation claims ownership of any part of the Moon's surface.[260] Russia, China, India, and the U.S. are party to the 1967 Outer Space Treaty,[261] which defines the Moon and all outer space as the "province of all mankind".[260] This treaty also restricts the use of the Moon to peaceful purposes, explicitly banning military installations and weapons of mass destruction.[262] The 1979 Moon Agreement was created to restrict the exploitation of the Moon's resources by any single nation, but as of January 2020, it has been signed and ratified by only 18 nations,[263] none of which engages in self-launched human space exploration. Although several individuals have made claims to the Moon in whole or in part, none of these are considered credible.[264][265][266]
	
	In 2020, U.S. President Donald Trump signed an executive order called "Encouraging International Support for the Recovery and Use of Space Resources". The order emphasizes that "the United States does not view outer space as a 'global commons'" and calls the Moon Agreement "a failed attempt at constraining free enterprise."[267][268]
	
	The Declaration of the Rights of the Moon[269] was created by a group of "lawyers, space archaeologists and concerned citizens" in 2021, drawing on precedents in the Rights of Nature movement and the concept of legal personality for non-human entities in space.[270]
	Coordination
	
	In light of future development on the Moon some international and multi-space agency organizations have been created:
	
		International Lunar Exploration Working Group (ILEWG)
		Moon Village Association (MVA)
		International Space Exploration Coordination Group (ISECG)
	
	In culture and life
	Luna, the Moon, from a 1550 edition of Guido Bonatti's Liber astronomiae
	See also: Moon in fiction and Tourism on the Moon
	Mythology
	Further information: Lunar deity, Selene, Luna (goddess), Man in the Moon, and Crescent
	Sun and Moon with faces (1493 woodcut)
	
	The contrast between the brighter highlands and the darker maria creates the patterns seen by different cultures as the Man in the Moon, the rabbit and the buffalo, among others. In many prehistoric and ancient cultures, the Moon was personified as a deity or other supernatural phenomenon, and astrological views of the Moon continue to be propagated.
	
	The ancient Sumerians believed that the Moon was the god Nanna,[271][272] who was the father of Inanna, the goddess of the planet Venus,[271][272] and Utu, the god of the Sun.[271][272] Nanna was later known as Sîn,[272][271] and was particularly associated with magic and sorcery.[271]
	
	In Mesopotamian iconography, the crescent was the primary symbol of Nanna-Sîn.[272] In ancient Greek art, the Moon goddess Selene was represented wearing a crescent on her headgear in an arrangement reminiscent of horns.[273][274] The star and crescent arrangement also goes back to the Bronze Age, representing either the Sun and Moon, or the Moon and planet Venus, in combination. It came to represent the goddess Artemis or Hecate, and via the patronage of Hecate came to be used as a symbol of Byzantium.
	
	An iconographic tradition of representing Sun and Moon with faces developed in the late medieval period.
	
	The splitting of the Moon (Arabic: انشقاق القمر‎) is a miracle attributed to Muhammad.[275] A song titled 'Moon Anthem' was released on the occasion of landing of India's Chandrayan-II on the Moon.[276]
	Calendar
	Further information: Lunar calendar, Lunisolar calendar, Metonic cycle, Blue moon, and Movable feast
	
	The Moon's regular phases make it a convenient timepiece, and the periods of its waxing and waning form the basis of many of the oldest calendars. Tally sticks, notched bones dating as far back as 20–30,000 years ago, are believed by some to mark the phases of the Moon.[277][278][279] The ~30-day month is an approximation of the lunar cycle. The English noun month and its cognates in other Germanic languages stem from Proto-Germanic *mǣnṓth-, which is connected to the above-mentioned Proto-Germanic *mǣnōn, indicating the usage of a lunar calendar among the Germanic peoples (Germanic calendar) prior to the adoption of a solar calendar.[280] The PIE root of moon, *méh1nōt, derives from the PIE verbal root *meh1-, "to measure", "indicat[ing] a functional conception of the Moon, i.e. marker of the month" (cf. the English words measure and menstrual),[281][282][283] and echoing the Moon's importance to many ancient cultures in measuring time (see Latin mensis and Ancient Greek μείς (meis) or μήν (mēn), meaning "month").[284][285][286][287] Most historical calendars are lunisolar. The 7th-century Islamic calendar is an example of a purely lunar calendar, where months are traditionally determined by the visual sighting of the hilal, or earliest crescent moon, over the horizon.[288]
	Lunar effect
	Main article: Lunar effect
	
	The lunar effect is a purported unproven correlation between specific stages of the roughly 29.5-day lunar cycle and behavior and physiological changes in living beings on Earth, including humans.
	
	The Moon has long been particularly associated with insanity and irrationality; the words lunacy and lunatic (popular shortening loony) are derived from the Latin name for the Moon, Luna. Philosophers Aristotle and Pliny the Elder argued that the full moon induced insanity in susceptible individuals, believing that the brain, which is mostly water, must be affected by the Moon and its power over the tides, but the Moon's gravity is too slight to affect any single person.[289] Even today, people who believe in a lunar effect claim that admissions to psychiatric hospitals, traffic accidents, homicides or suicides increase during a full moon, but dozens of studies invalidate these claims.[289][290][291][292][293]
	See also
	
		List of natural satellites
	
	Notes
	
	Between 18.29° and 28.58° to Earth's equator.[1]
	There are a number of near-Earth asteroids, including 3753 Cruithne, that are co-orbital with Earth: their orbits bring them close to Earth for periods of time but then alter in the long term (Morais et al, 2002). These are quasi-satellites – they are not moons as they do not orbit Earth. For more information, see Other moons of Earth.
	The maximum value is given based on scaling of the brightness from the value of −12.74 given for an equator to Moon-centre distance of 378 000 km in the NASA factsheet reference to the minimum Earth–Moon distance given there, after the latter is corrected for Earth's equatorial radius of 6 378 km, giving 350 600 km. The minimum value (for a distant new moon) is based on a similar scaling using the maximum Earth–Moon distance of 407 000 km (given in the factsheet) and by calculating the brightness of the earthshine onto such a new moon. The brightness of the earthshine is [ Earth albedo × (Earth radius / Radius of Moon's orbit)2 ] relative to the direct solar illumination that occurs for a full moon. (Earth albedo = 0.367; Earth radius = (polar radius × equatorial radius)½ = 6 367 km.)
	The range of angular size values given are based on simple scaling of the following values given in the fact sheet reference: at an Earth-equator to Moon-centre distance of 378 000 km, the angular size is 1896 arcseconds. The same fact sheet gives extreme Earth–Moon distances of 407 000 km and 357 000 km. For the maximum angular size, the minimum distance has to be corrected for Earth's equatorial radius of 6 378 km, giving 350 600 km.
	Lucey et al. (2006) give 107 particles cm−3 by day and 105 particles cm−3 by night. Along with equatorial surface temperatures of 390 K by day and 100 K by night, the ideal gas law yields the pressures given in the infobox (rounded to the nearest order of magnitude): 10−7 Pa by day and 10−10 Pa by night.
	Charon is larger with respect to Pluto, but Pluto is a dwarf planet.
	More accurately, the Moon's mean sidereal period (fixed star to fixed star) is 27.321661 days (27 d 07 h 43 min 11.5 s), and its mean tropical orbital period (from equinox to equinox) is 27.321582 days (27 d 07 h 43 min 04.7 s) (Explanatory Supplement to the Astronomical Ephemeris, 1961, at p.107).
	More accurately, the Moon's mean synodic period (between mean solar conjunctions) is 29.530589 days (29 d 12 h 44 min 02.9 s) (Explanatory Supplement to the Astronomical Ephemeris, 1961, at p.107).
	There is no strong correlation between the sizes of planets and the sizes of their satellites. Larger planets tend to have more satellites, both large and small, than smaller planets.
	With 27% the diameter and 60% the density of Earth, the Moon has 1.23% of the mass of Earth. The moon Charon is larger relative to its primary Pluto, but Pluto is now considered to be a dwarf planet.
	The Sun's apparent magnitude is −26.7, while the full moon's apparent magnitude is −12.7.
	See graph in Sun#Life phases. At present, the diameter of the Sun is increasing at a rate of about five percent per billion years. This is very similar to the rate at which the apparent angular diameter of the Moon is decreasing as it recedes from Earth.
	
		On average, the Moon covers an area of 0.21078 square degrees on the night sky.
	
	References
	Citations
	
	Wieczorek, Mark A.; Jolliff, Bradley L.; Khan, Amir; Pritchard, Matthew E.; Weiss, Benjamin P.; Williams, James G.; Hood, Lon L.; Righter, Kevin; Neal, Clive R.; Shearer, Charles K.; McCallum, I. Stewart; Tompkins, Stephanie; Hawke, B. Ray; Peterson, Chris; Gillis, Jeffrey J.; Bussey, Ben (2006). "The constitution and structure of the lunar interior". Reviews in Mineralogy and Geochemistry. 60 (1): 221–364. Bibcode:2006RvMG...60..221W. doi:10.2138/rmg.2006.60.3. S2CID 130734866. Archived from the original on 19 August 2020. Retrieved 2 December 2019.
	Lang, Kenneth R. (2011). The Cambridge Guide to the Solar System (2nd ed.). Cambridge University Press. ISBN 9781139494175. Archived from the original on 1 January 2016.
	Morais, M. H. M.; Morbidelli, A. (2002). "The Population of Near-Earth Asteroids in Coorbital Motion with the Earth". Icarus. 160 (1): 1–9. Bibcode:2002Icar..160....1M. doi:10.1006/icar.2002.6937. hdl:10316/4391. S2CID 55214551. Archived from the original on 19 August 2020. Retrieved 2 December 2019.
	Williams, David R. (2 February 2006). "Moon Fact Sheet". NASA/National Space Science Data Center. Archived from the original on 23 March 2010. Retrieved 31 December 2008.
	Smith, David E.; Zuber, Maria T.; Neumann, Gregory A.; Lemoine, Frank G. (1 January 1997). "Topography of the Moon from the Clementine lidar". Journal of Geophysical Research. 102 (E1): 1601. Bibcode:1997JGR...102.1591S. doi:10.1029/96JE02940. hdl:2060/19980018849. S2CID 17475023. Archived from the original on 19 August 2020. Retrieved 2 December 2019.
	Terry, Paul (2013). Top 10 of Everything. Octopus Publishing Group Ltd. p. 226. ISBN 978-0-600-62887-3.
	Williams, James G.; Newhall, XX; Dickey, Jean O. (1996). "Lunar moments, tides, orientation, and coordinate frames". Planetary and Space Science. 44 (10): 1077–1080. Bibcode:1996P&SS...44.1077W. doi:10.1016/0032-0633(95)00154-9.
	Hamilton, Calvin J.; Hamilton, Rosanna L., The Moon, Views of the Solar System Archived 4 February 2016 at the Wayback Machine, 1995–2011.
	Makemson, Maud W. (1971). "Determination of selenographic positions". The Moon. 2 (3): 293–308. Bibcode:1971Moon....2..293M. doi:10.1007/BF00561882. S2CID 119603394.
	Archinal, Brent A.; A'Hearn, Michael F.; Bowell, Edward G.; Conrad, Albert R.; Consolmagno, Guy J.; Courtin, Régis; Fukushima, Toshio; Hestroffer, Daniel; Hilton, James L.; Krasinsky, George A.; Neumann, Gregory A.; Oberst, Jürgen; Seidelmann, P. Kenneth; Stooke, Philip J.; Tholen, David J.; Thomas, Paul C.; Williams, Iwan P. (2010). "Report of the IAU Working Group on Cartographic Coordinates and Rotational Elements: 2009" (PDF). Celestial Mechanics and Dynamical Astronomy. 109 (2): 101–135. Bibcode:2011CeMDA.109..101A. doi:10.1007/s10569-010-9320-4. S2CID 189842666. Archived from the original (PDF) on 4 March 2016. Retrieved 24 September 2018. also available "via usgs.gov" (PDF). Archived (PDF) from the original on 27 April 2019. Retrieved 26 September 2018.
	Matthews, Grant (2008). "Celestial body irradiance determination from an underfilled satellite radiometer: application to albedo and thermal emission measurements of the Moon using CERES". Applied Optics. 47 (27): 4981–4993. Bibcode:2008ApOpt..47.4981M. doi:10.1364/AO.47.004981. PMID 18806861.
	Bugby, D. C.; Farmer, J. T.; O’Connor, B. F.; Wirzburger, M. J.; C. J. Stouffer, E. D. Abel (January 2010). Two‐Phase Thermal Switching System for a Small, Extended Duration Lunar Surface Science Platform. AIP Conference Proceedings. 1208. pp. 76–83. Bibcode:2010AIPC.1208...76B. doi:10.1063/1.3326291. hdl:2060/20100009810.
	Vasavada, A. R.; Paige, D. A.; Wood, S. E. (1999). "Near-Surface Temperatures on Mercury and the Moon and the Stability of Polar Ice Deposits". Icarus. 141 (2): 179–193. Bibcode:1999Icar..141..179V. doi:10.1006/icar.1999.6175. S2CID 37706412. Archived from the original on 19 August 2020. Retrieved 2 December 2019.
	Lucey, Paul; Korotev, Randy L.; Gillis, Jeffrey J.; Taylor, Larry A.; Lawrence, David; Campbell, Bruce A.; Elphic, Rick; Feldman, Bill; Hood, Lon L.; Hunten, Donald; Mendillo, Michael; Noble, Sarah; Papike, James J.; Reedy, Robert C.; Lawson, Stefanie; Prettyman, Tom; Gasnault, Olivier; Maurice, Sylvestre (2006). "Understanding the lunar surface and space-Moon interactions". Reviews in Mineralogy and Geochemistry. 60 (1): 83–219. Bibcode:2006RvMG...60...83L. doi:10.2138/rmg.2006.60.2.
	Horner, Jonti (18 July 2019). "How big is the Moon?". Archived from the original on 7 November 2020. Retrieved 15 November 2020.
	"By the Numbers | Earth's Moon". NASA Solar System Exploration. NASA. Retrieved 15 December 2020.
	Stern, David (30 March 2014). "Libration of the Moon". NASA. Archived from the original on 22 May 2020. Retrieved 11 February 2020.
	"Naming Astronomical Objects: Spelling of Names". International Astronomical Union. Archived from the original on 16 December 2008. Retrieved 6 April 2020.
	"Gazetteer of Planetary Nomenclature: Planetary Nomenclature FAQ". USGS Astrogeology Research Program. Archived from the original on 27 May 2010. Retrieved 6 April 2020.
	Orel, Vladimir (2003). A Handbook of Germanic Etymology. Brill. Archived from the original on 17 June 2020. Retrieved 5 March 2020.
	López-Menchero, Fernando (22 May 2020). "Late Proto-Indo-European Etymological Lexicon".
	Barnhart, Robert K. (1995). The Barnhart Concise Dictionary of Etymology. Harper Collins. p. 487. ISBN 978-0-06-270084-1.
	E.g.: Hall III, James A. (2016). Moons of the Solar System. Springer International. ISBN 978-3-319-20636-3.
	"Luna". Oxford English Dictionary (Online ed.). Oxford University Press. (Subscription or participating institution membership required.)
	"Cynthia". Oxford English Dictionary (Online ed.). Oxford University Press. (Subscription or participating institution membership required.)
	"selenian". Merriam-Webster Dictionary.
	"selenian". Oxford English Dictionary (Online ed.). Oxford University Press. (Subscription or participating institution membership required.)
	"selenic". Oxford English Dictionary (Online ed.). Oxford University Press. (Subscription or participating institution membership required.)
	"selenic". Merriam-Webster Dictionary.
	"Oxford English Dictionary: lunar, a. and n." Oxford English Dictionary: Second Edition 1989. Oxford University Press. Archived from the original on 19 August 2020. Retrieved 23 March 2010.
	σελήνη. Liddell, Henry George; Scott, Robert; A Greek–English Lexicon at the Perseus Project.
	Pannen, Imke (2010). When the Bad Bleeds: Mantic Elements in English Renaissance Revenge Tragedy. V&R unipress GmbH. pp. 96–. ISBN 978-3-89971-640-5. Archived from the original on 4 September 2016.
	Thiemens, Maxwell M.; Sprung, Peter; Fonseca, Raúl O. C.; Leitzke, Felipe P.; Münker, Carsten (July 2019). "Early Moon formation inferred from hafnium-tungsten systematics". Nature Geoscience. 12 (9): 696–700. Bibcode:2019NatGe..12..696T. doi:10.1038/s41561-019-0398-3. S2CID 198997377.
	"The Moon is older than scientists thought". Universe Today. Archived from the original on 3 August 2019. Retrieved 3 August 2019.
	Barboni, M.; Boehnke, P.; Keller, C.B.; Kohl, I.E.; Schoene, B.; Young, E.D.; McKeegan, K.D. (2017). "Early formation of the Moon 4.51 billion years ago". Science Advances. 3 (1): e1602365. Bibcode:2017SciA....3E2365B. doi:10.1126/sciadv.1602365. PMC 5226643. PMID 28097222.
	Binder, A. B. (1974). "On the origin of the Moon by rotational fission". The Moon. 11 (2): 53–76. Bibcode:1974Moon...11...53B. doi:10.1007/BF01877794. S2CID 122622374.
	Stroud, Rick (2009). The Book of the Moon. Walken and Company. pp. 24–27. ISBN 978-0-8027-1734-4. Archived from the original on 17 June 2020. Retrieved 11 November 2019.
	Mitler, H. E. (1975). "Formation of an iron-poor moon by partial capture, or: Yet another exotic theory of lunar origin". Icarus. 24 (2): 256–268. Bibcode:1975Icar...24..256M. doi:10.1016/0019-1035(75)90102-5.
	Stevenson, D.J. (1987). "Origin of the moon–The collision hypothesis". Annual Review of Earth and Planetary Sciences. 15 (1): 271–315. Bibcode:1987AREPS..15..271S. doi:10.1146/annurev.ea.15.050187.001415. S2CID 53516498. Archived from the original on 19 August 2020. Retrieved 2 December 2019.
	Taylor, G. Jeffrey (31 December 1998). "Origin of the Earth and Moon". Planetary Science Research Discoveries. Hawai'i Institute of Geophysics and Planetology. Archived from the original on 10 June 2010. Retrieved 7 April 2010.
	"Asteroids Bear Scars of Moon's Violent Formation". 16 April 2015. Archived from the original on 8 October 2016.
	van Putten, Maurice H. P. M. (July 2017). "Scaling in global tidal dissipation of the Earth-Moon system". New Astronomy. 54: 115–121. arXiv:1609.07474. Bibcode:2017NewA...54..115V. doi:10.1016/j.newast.2017.01.012. S2CID 119285032.
	Canup, R.; Asphaug, E. (2001). "Origin of the Moon in a giant impact near the end of Earth's formation". Nature. 412 (6848): 708–712. Bibcode:2001Natur.412..708C. doi:10.1038/35089010. PMID 11507633. S2CID 4413525.
	"Earth-Asteroid Collision Formed Moon Later Than Thought". National Geographic. 28 October 2010. Archived from the original on 18 April 2009. Retrieved 7 May 2012.
	Kleine, Thorsten (2008). "2008 Pellas-Ryder Award for Mathieu Touboul" (PDF). Meteoritics and Planetary Science. 43 (S7): A11–A12. Bibcode:2008M&PS...43...11K. doi:10.1111/j.1945-5100.2008.tb00709.x. Archived from the original (PDF) on 27 July 2018. Retrieved 8 April 2020.
	Touboul, M.; Kleine, T.; Bourdon, B.; Palme, H.; Wieler, R. (2007). "Late formation and prolonged differentiation of the Moon inferred from W isotopes in lunar metals". Nature. 450 (7173): 1206–1209. Bibcode:2007Natur.450.1206T. doi:10.1038/nature06428. PMID 18097403. S2CID 4416259.
	"Flying Oceans of Magma Help Demystify the Moon's Creation". National Geographic. 8 April 2015. Archived from the original on 9 April 2015.
	Pahlevan, Kaveh; Stevenson, David J. (2007). "Equilibration in the aftermath of the lunar-forming giant impact". Earth and Planetary Science Letters. 262 (3–4): 438–449. arXiv:1012.5323. Bibcode:2007E&PSL.262..438P. doi:10.1016/j.epsl.2007.07.055. S2CID 53064179.
	Nield, Ted (2009). "Moonwalk (summary of meeting at Meteoritical Society's 72nd Annual Meeting, Nancy, France)". Geoscientist. Vol. 19. p. 8. Archived from the original on 27 September 2012.
	Warren, P. H. (1985). "The magma ocean concept and lunar evolution". Annual Review of Earth and Planetary Sciences. 13 (1): 201–240. Bibcode:1985AREPS..13..201W. doi:10.1146/annurev.ea.13.050185.001221.
	Tonks, W. Brian; Melosh, H. Jay (1993). "Magma ocean formation due to giant impacts". Journal of Geophysical Research. 98 (E3): 5319–5333. Bibcode:1993JGR....98.5319T. doi:10.1029/92JE02726.
	Daniel Clery (11 October 2013). "Impact Theory Gets Whacked". Science. 342 (6155): 183–185. Bibcode:2013Sci...342..183C. doi:10.1126/science.342.6155.183. PMID 24115419.
	Wiechert, U.; Halliday, A. N.; Lee, D.-C.; Snyder, G. A.; Taylor, L. A.; Rumble, D. (October 2001). "Oxygen Isotopes and the Moon-Forming Giant Impact". Science. 294 (12): 345–348. Bibcode:2001Sci...294..345W. doi:10.1126/science.1063037. PMID 11598294. S2CID 29835446. Archived from the original on 20 April 2009. Retrieved 5 July 2009.
	Pahlevan, Kaveh; Stevenson, David (October 2007). "Equilibration in the Aftermath of the Lunar-forming Giant Impact". Earth and Planetary Science Letters. 262 (3–4): 438–449. arXiv:1012.5323. Bibcode:2007E&PSL.262..438P. doi:10.1016/j.epsl.2007.07.055. S2CID 53064179.
	"Titanium Paternity Test Says Earth is the Moon's Only Parent (University of Chicago)". Astrobio.net. 5 April 2012. Archived from the original on 8 August 2012. Retrieved 3 October 2013.
	Garrick-Bethell, Ian; Perera, Viranga; Nimmo, Francis; Zuber, Maria T. (2014). "The tidal-rotational shape of the Moon and evidence for polar wander" (PDF). Nature. 512 (7513): 181–184. Bibcode:2014Natur.512..181G. doi:10.1038/nature13639. PMID 25079322. S2CID 4452886. Archived (PDF) from the original on 4 August 2020. Retrieved 12 April 2020.
	Taylor, Stuart R. (1975). Lunar Science: a Post-Apollo View. Oxford: Pergamon Press. p. 64. ISBN 978-0-08-018274-2.
	Runcorn, Stanley Keith (31 March 1977). "Interpretation of lunar potential fields". Philosophical Transactions of the Royal Society of London. Series A, Mathematical and Physical Sciences. 285 (1327): 507–516. Bibcode:1977RSPTA.285..507R. doi:10.1098/rsta.1977.0094. S2CID 124703189.
	Brown, D.; Anderson, J. (6 January 2011). "NASA Research Team Reveals Moon Has Earth-Like Core". NASA. NASA. Archived from the original on 11 January 2012.
	Weber, R.C.; Lin, P.-Y.; Garnero, E.J.; Williams, Q.; Lognonne, P. (21 January 2011). "Seismic Detection of the Lunar Core" (PDF). Science. 331 (6015): 309–312. Bibcode:2011Sci...331..309W. doi:10.1126/science.1199375. PMID 21212323. S2CID 206530647. Archived from the original (PDF) on 15 October 2015. Retrieved 10 April 2017.
	Nemchin, A.; Timms, N.; Pidgeon, R.; Geisler, T.; Reddy, S.; Meyer, C. (2009). "Timing of crystallization of the lunar magma ocean constrained by the oldest zircon". Nature Geoscience. 2 (2): 133–136. Bibcode:2009NatGe...2..133N. doi:10.1038/ngeo417. hdl:20.500.11937/44375.
	Shearer, Charles K.; Hess, Paul C.; Wieczorek, Mark A.; Pritchard, Matt E.; Parmentier, E. Mark; Borg, Lars E.; Longhi, John; Elkins-Tanton, Linda T.; Neal, Clive R.; Antonenko, Irene; Canup, Robin M.; Halliday, Alex N.; Grove, Tim L.; Hager, Bradford H.; Lee, D.-C.; Wiechert, Uwe (2006). "Thermal and magmatic evolution of the Moon". Reviews in Mineralogy and Geochemistry. 60 (1): 365–518. Bibcode:2006RvMG...60..365S. doi:10.2138/rmg.2006.60.4. S2CID 129184748. Archived from the original on 19 August 2020. Retrieved 2 December 2019.
	Schubert, J. (2004). "Interior composition, structure, and dynamics of the Galilean satellites.". In F. Bagenal; et al. (eds.). Jupiter: The Planet, Satellites, and Magnetosphere. Cambridge University Press. pp. 281–306. ISBN 978-0-521-81808-7.
	Williams, J.G.; Turyshev, S.G.; Boggs, D.H.; Ratcliff, J.T. (2006). "Lunar laser ranging science: Gravitational physics and lunar interior and geodesy". Advances in Space Research. 37 (1): 67–71. arXiv:gr-qc/0412049. Bibcode:2006AdSpR..37...67W. doi:10.1016/j.asr.2005.05.013. S2CID 14801321.
	Evans, Alexander J.; Tikoo, Sonia M.; Jeffrey C., Andrews-Hanna (January 2018). "The Case Against an Early Lunar Dynamo Powered by Core Convection". Geophysical Research Letters. 45 (1): 98–107. Bibcode:2018GeoRL..45...98E. doi:10.1002/2017GL075441.
	Mighani, S.; Wang, H.; Shuster, D.L.; Borlina, C.S.; Nichols, C.I.O.; Weiss, B.P. (2020). "The end of the lunar dynamo". Science Advances. 6 (1): eaax0883. Bibcode:2020SciA....6..883M. doi:10.1126/sciadv.aax0883. PMC 6938704. PMID 31911941.
	Garrick-Bethell, Ian; Weiss, iBenjamin P.; Shuster, David L.; Buz, Jennifer (2009). "Early Lunar Magnetism". Science. 323 (5912): 356–359. Bibcode:2009Sci...323..356G. doi:10.1126/science.1166804. PMID 19150839. S2CID 23227936. Archived from the original on 19 August 2020. Retrieved 2 December 2019.
	"Magnetometer / Electron Reflectometer Results". Lunar Prospector (NASA). 2001. Archived from the original on 27 May 2010. Retrieved 17 March 2010.
	Hood, L.L.; Huang, Z. (1991). "Formation of magnetic anomalies antipodal to lunar impact basins: Two-dimensional model calculations". Journal of Geophysical Research. 96 (B6): 9837–9846. Bibcode:1991JGR....96.9837H. doi:10.1029/91JB00308.
	Spudis, Paul D.; Cook, A.; Robinson, M.; Bussey, B.; Fessler, B. (January 1998). "Topography of the South Polar Region from Clementine Stereo Imaging". Workshop on New Views of the Moon: Integrated Remotely Sensed, Geophysical, and Sample Datasets: 69. Bibcode:1998nvmi.conf...69S.
	Spudis, Paul D.; Reisse, Robert A.; Gillis, Jeffrey J. (1994). "Ancient Multiring Basins on the Moon Revealed by Clementine Laser Altimetry". Science. 266 (5192): 1848–1851. Bibcode:1994Sci...266.1848S. doi:10.1126/science.266.5192.1848. PMID 17737079. S2CID 41861312.
	Pieters, C. M.; Tompkins, S.; Head, J. W.; Hess, P. C. (1997). "Mineralogy of the Mafic Anomaly in the South Pole‐Aitken Basin: Implications for excavation of the lunar mantle". Geophysical Research Letters. 24 (15): 1903–1906. Bibcode:1997GeoRL..24.1903P. doi:10.1029/97GL01718. hdl:2060/19980018038.
	Taylor, G. J. (17 July 1998). "The Biggest Hole in the Solar System". Planetary Science Research Discoveries: 20. Bibcode:1998psrd.reptE..20T. Archived from the original on 20 August 2007. Retrieved 12 April 2007.
	Schultz, P.H. (March 1997). "Forming the south-pole Aitken basin – The extreme games". Conference Paper, 28th Annual Lunar and Planetary Science Conference. 28: 1259. Bibcode:1997LPI....28.1259S.
	"NASA's LRO Reveals 'Incredible Shrinking Moon'". NASA. 19 August 2010. Archived from the original on 21 August 2010.
	Watters, Thomas R.; Weber, Renee C.; Collins, Geoffrey C.; Howley, Ian J.; Schmerr, Nicholas C.; Johnson, Catherine L. (June 2019). "Shallow seismic activity and young thrust faults on the Moon". Nature Geoscience (published 13 May 2019). 12 (6): 411–417. Bibcode:2019NatGe..12..411W. doi:10.1038/s41561-019-0362-2. ISSN 1752-0894. S2CID 182137223.
	Wlasuk, Peter (2000). Observing the Moon. Springer. p. 19. ISBN 978-1-85233-193-1.
	Norman, M. (21 April 2004). "The Oldest Moon Rocks". Planetary Science Research Discoveries. Hawai'i Institute of Geophysics and Planetology. Archived from the original on 18 April 2007. Retrieved 12 April 2007.
	Wilson, Lionel; Head, James W. (2003). "Lunar Gruithuisen and Mairan domes: Rheology and mode of emplacement". Journal of Geophysical Research. 108 (E2): 5012. Bibcode:2003JGRE..108.5012W. CiteSeerX 10.1.1.654.9619. doi:10.1029/2002JE001909. Archived from the original on 12 March 2007. Retrieved 12 April 2007.
	Spudis, P. D. (2004). "Moon". World Book Online Reference Center, NASA. Archived from the original on 3 July 2013. Retrieved 12 April 2007.
	Gillis, J. J.; Spudis, P. D. (1996). "The Composition and Geologic Setting of Lunar Far Side Maria". Lunar and Planetary Science. 27: 413. Bibcode:1996LPI....27..413G.
	Lawrence, D. J.; Feldman, W. C.; Barraclough, B. L.; Binder, A. B.; Elphic, R. C.; Maurice, S.; Thomsen, D. R. (11 August 1998). "Global Elemental Maps of the Moon: The Lunar Prospector Gamma-Ray Spectrometer". Science. 281 (5382): 1484–1489. Bibcode:1998Sci...281.1484L. doi:10.1126/science.281.5382.1484. PMID 9727970.
	Taylor, G. J. (31 August 2000). "A New Moon for the Twenty-First Century". Planetary Science Research Discoveries: 41. Bibcode:2000psrd.reptE..41T. Archived from the original on 1 March 2012. Retrieved 12 April 2007.
	Papike, J.; Ryder, G.; Shearer, C. (1998). "Lunar Samples". Reviews in Mineralogy and Geochemistry. 36: 5.1–5.234.
	Hiesinger, H.; Head, J. W.; Wolf, U.; Jaumann, R.; Neukum, G. (2003). "Ages and stratigraphy of mare basalts in Oceanus Procellarum, Mare Numbium, Mare Cognitum, and Mare Insularum". Journal of Geophysical Research. 108 (E7): 1029. Bibcode:2003JGRE..108.5065H. doi:10.1029/2002JE001985. S2CID 9570915.
	Phil Berardelli (9 November 2006). "Long Live the Moon!". Science. Archived from the original on 18 October 2014. Retrieved 14 October 2014.
	Jason Major (14 October 2014). "Volcanoes Erupted 'Recently' on the Moon". Discovery News. Archived from the original on 16 October 2014.
	"NASA Mission Finds Widespread Evidence of Young Lunar Volcanism". NASA. 12 October 2014. Archived from the original on 3 January 2015.
	Eric Hand (12 October 2014). "Recent volcanic eruptions on the moon". Science. Archived from the original on 14 October 2014.
	Braden, S.E.; Stopar, J.D.; Robinson, M.S.; Lawrence, S.J.; van der Bogert, C.H.; Hiesinger, H. (2014). "Evidence for basaltic volcanism on the Moon within the past 100 million years". Nature Geoscience. 7 (11): 787–791. Bibcode:2014NatGe...7..787B. doi:10.1038/ngeo2252.
	Srivastava, N.; Gupta, R.P. (2013). "Young viscous flows in the Lowell crater of Orientale basin, Moon: Impact melts or volcanic eruptions?". Planetary and Space Science. 87: 37–45. Bibcode:2013P&SS...87...37S. doi:10.1016/j.pss.2013.09.001.
	Gupta, R.P.; Srivastava, N.; Tiwari, R.K. (2014). "Evidences of relatively new volcanic flows on the Moon". Current Science. 107 (3): 454–460. JSTOR 24103498.
	Whitten, Jennifer; Head, James W.; Staid, Matthew; Pieters, Carle M.; Mustard, John; Clark, Roger; Nettles, Jeff; Klima, Rachel L.; Taylor, Larry (2011). "Lunar mare deposits associated with the Orientale impact basin: New insights into mineralogy, history, mode of emplacement, and relation to Orientale Basin evolution from Moon Mineralogy Mapper (M3) data from Chandrayaan-1". Journal of Geophysical Research. 116: E00G09. Bibcode:2011JGRE..116.0G09W. doi:10.1029/2010JE003736. S2CID 7234547.
	Cho, Y.; et al. (2012). "Young mare volcanism in the Orientale region contemporary with the Procellarum KREEP Terrane (PKT) volcanism peak period 2 b.y. ago". Geophysical Research Letters. 39 (11): L11203. Bibcode:2012GeoRL..3911203C. doi:10.1029/2012GL051838.
	Munsell, K. (4 December 2006). "Majestic Mountains". Solar System Exploration. NASA. Archived from the original on 17 September 2008. Retrieved 12 April 2007.
	Richard Lovett (2011). "Early Earth may have had two moons : Nature News". Nature. doi:10.1038/news.2011.456. Archived from the original on 3 November 2012. Retrieved 1 November 2012.
	"Was our two-faced moon in a small collision?". Theconversation.edu.au. Archived from the original on 30 January 2013. Retrieved 1 November 2012.
	Quillen, Alice C.; Martini, Larkin; Nakajima, Miki (September 2019). "Near/far side asymmetry in the tidally heated Moon". Icarus. 329: 182–196. arXiv:1810.10676. Bibcode:2019Icar..329..182Q. doi:10.1016/j.icarus.2019.04.010. PMC 7489467. PMID 32934397.
	Melosh, H. J. (1989). Impact cratering: A geologic process. Oxford University Press. ISBN 978-0-19-504284-9.
	"Moon Facts". SMART-1. European Space Agency. 2010. Archived from the original on 17 March 2012. Retrieved 12 May 2010.
	Wilhelms, Don (1987). "Relative Ages" (PDF). Geologic History of the Moon. U.S. Geological Survey. Archived from the original (PDF) on 11 June 2010. Retrieved 4 April 2010.
	Hartmann, William K.; Quantin, Cathy; Mangold, Nicolas (2007). "Possible long-term decline in impact rates: 2. Lunar impact-melt data regarding impact history". Icarus. 186 (1): 11–23. Bibcode:2007Icar..186...11H. doi:10.1016/j.icarus.2006.09.009.
	"The Smell of Moondust". NASA. 30 January 2006. Archived from the original on 8 March 2010. Retrieved 15 March 2010.
	Heiken, G. (1991). Vaniman, D.; French, B. (eds.). Lunar Sourcebook, a user's guide to the Moon. New York: Cambridge University Press. p. 736. ISBN 978-0-521-33444-0. Archived from the original on 17 June 2020. Retrieved 17 December 2019.
	Rasmussen, K.L.; Warren, P.H. (1985). "Megaregolith thickness, heat flow, and the bulk composition of the Moon". Nature. 313 (5998): 121–124. Bibcode:1985Natur.313..121R. doi:10.1038/313121a0. S2CID 4245137.
	Boyle, Rebecca. "The moon has hundreds more craters than we thought". Archived from the original on 13 October 2016.
	Speyerer, Emerson J.; Povilaitis, Reinhold Z.; Robinson, Mark S.; Thomas, Peter C.; Wagner, Robert V. (13 October 2016). "Quantifying crater production and regolith overturn on the Moon with temporal imaging". Nature. 538 (7624): 215–218. Bibcode:2016Natur.538..215S. doi:10.1038/nature19829. PMID 27734864. S2CID 4443574.
	"Earth's Moon Hit by Surprising Number of Meteoroids". NASA. 13 October 2016. Retrieved 21 May 2021.
	Muller, P.; Sjogren, W. (1968). "Mascons: lunar mass concentrations". Science. 161 (3842): 680–684. Bibcode:1968Sci...161..680M. doi:10.1126/science.161.3842.680. PMID 17801458. S2CID 40110502.
	Richard A. Kerr (12 April 2013). "The Mystery of Our Moon's Gravitational Bumps Solved?". Science. 340 (6129): 138–139. doi:10.1126/science.340.6129.138-a. PMID 23580504.
	Konopliv, A.; Asmar, S.; Carranza, E.; Sjogren, W.; Yuan, D. (2001). "Recent gravity models as a result of the Lunar Prospector mission" (PDF). Icarus. 50 (1): 1–18. Bibcode:2001Icar..150....1K. CiteSeerX 10.1.1.18.1930. doi:10.1006/icar.2000.6573. Archived from the original (PDF) on 13 November 2004.
	Chrbolková, Kateřina; Kohout, Tomáš; Ďurech, Josef (November 2019). "Reflectance spectra of seven lunar swirls examined by statistical methods: A space weathering study". Icarus. 333: 516–527. Bibcode:2019Icar..333..516C. doi:10.1016/j.icarus.2019.05.024.
	Margot, J. L.; Campbell, D. B.; Jurgens, R. F.; Slade, M. A. (4 June 1999). "Topography of the Lunar Poles from Radar Interferometry: A Survey of Cold Trap Locations" (PDF). Science. 284 (5420): 1658–1660. Bibcode:1999Sci...284.1658M. CiteSeerX 10.1.1.485.312. doi:10.1126/science.284.5420.1658. PMID 10356393. Archived (PDF) from the original on 11 August 2017. Retrieved 25 October 2017.
	Ward, William R. (1 August 1975). "Past Orientation of the Lunar Spin Axis". Science. 189 (4200): 377–379. Bibcode:1975Sci...189..377W. doi:10.1126/science.189.4200.377. PMID 17840827. S2CID 21185695.
	Martel, L. M. V. (4 June 2003). "The Moon's Dark, Icy Poles". Planetary Science Research Discoveries: 73. Bibcode:2003psrd.reptE..73M. Archived from the original on 1 March 2012. Retrieved 12 April 2007.
	Seedhouse, Erik (2009). Lunar Outpost: The Challenges of Establishing a Human Settlement on the Moon. Springer-Praxis Books in Space Exploration. Germany: Springer Praxis. p. 136. ISBN 978-0-387-09746-6. Archived from the original on 26 November 2020. Retrieved 22 August 2020.
	Coulter, Dauna (18 March 2010). "The Multiplying Mystery of Moonwater". NASA. Archived from the original on 13 December 2012. Retrieved 28 March 2010.
	Spudis, P. (6 November 2006). "Ice on the Moon". The Space Review. Archived from the original on 22 February 2007. Retrieved 12 April 2007.
	Feldman, W. C.; Maurice, S.; Binder, A. B.; Barraclough, B. L.; R.C. Elphic; D.J. Lawrence (1998). "Fluxes of Fast and Epithermal Neutrons from Lunar Prospector: Evidence for Water Ice at the Lunar Poles". Science. 281 (5382): 1496–1500. Bibcode:1998Sci...281.1496F. doi:10.1126/science.281.5382.1496. PMID 9727973. S2CID 9005608.
	Saal, Alberto E.; Hauri, Erik H.; Cascio, Mauro L.; van Orman, James A.; Rutherford, Malcolm C.; Cooper, Reid F. (2008). "Volatile content of lunar volcanic glasses and the presence of water in the Moon's interior". Nature. 454 (7201): 192–195. Bibcode:2008Natur.454..192S. doi:10.1038/nature07047. PMID 18615079. S2CID 4394004.
	Pieters, C. M.; Goswami, J. N.; Clark, R. N.; Annadurai, M.; Boardman, J.; Buratti, B.; Combe, J.-P.; Dyar, M. D.; Green, R.; Head, J. W.; Hibbitts, C.; Hicks, M.; Isaacson, P.; Klima, R.; Kramer, G.; Kumar, S.; Livo, E.; Lundeen, S.; Malaret, E.; McCord, T.; Mustard, J.; Nettles, J.; Petro, N.; Runyon, C.; Staid, M.; Sunshine, J.; Taylor, L.A.; Tompkins, S.; Varanasi, P. (2009). "Character and Spatial Distribution of OH/H2O on the Surface of the Moon Seen by M3 on Chandrayaan-1". Science. 326 (5952): 568–572. Bibcode:2009Sci...326..568P. doi:10.1126/science.1178658. PMID 19779151. S2CID 447133. Archived from the original on 19 August 2020. Retrieved 2 December 2019.
	Li, Shuai; Lucey, Paul G.; Milliken, Ralph E.; Hayne, Paul O.; Fisher, Elizabeth; Williams, Jean-Pierre; Hurley, Dana M.; Elphic, Richard C. (August 2018). "Direct evidence of surface exposed water ice in the lunar polar regions". Proceedings of the National Academy of Sciences. 115 (36): 8907–8912. Bibcode:2018PNAS..115.8907L. doi:10.1073/pnas.1802345115. PMC 6130389. PMID 30126996.
	Lakdawalla, Emily (13 November 2009). "LCROSS Lunar Impactor Mission: "Yes, We Found Water!"". The Planetary Society. Archived from the original on 22 January 2010. Retrieved 13 April 2010.
	Colaprete, A.; Ennico, K.; Wooden, D.; Shirley, M.; Heldmann, J.; Marshall, W.; Sollitt, L.; Asphaug, E.; Korycansky, D.; Schultz, P.; Hermalyn, B.; Galal, K.; Bart, G.D.; Goldstein, D.; Summy, D. (1–5 March 2010). "Water and More: An Overview of LCROSS Impact Results". 41st Lunar and Planetary Science Conference. 41 (1533): 2335. Bibcode:2010LPI....41.2335C.
	Colaprete, Anthony; Schultz, Peter; Heldmann, Jennifer; Wooden, Diane; Shirley, Mark; Ennico, Kimberly; Hermalyn, Brendan; Marshall, William; Ricco, Antonio; Elphic, Richard C.; Goldstein, David; Summy, Dustin; Bart, Gwendolyn D.; Asphaug, Erik; Korycansky, Don; Landis, David; Sollitt, Luke (22 October 2010). "Detection of Water in the LCROSS Ejecta Plume". Science. 330 (6003): 463–468. Bibcode:2010Sci...330..463C. doi:10.1126/science.1186986. PMID 20966242. S2CID 206525375. Archived from the original on 19 August 2020. Retrieved 2 December 2019.
	Hauri, Erik; Thomas Weinreich; Albert E. Saal; Malcolm C. Rutherford; James A. Van Orman (26 May 2011). "High Pre-Eruptive Water Contents Preserved in Lunar Melt Inclusions". Science Express. 10 (1126): 213–215. Bibcode:2011Sci...333..213H. doi:10.1126/science.1204626. PMID 21617039. S2CID 44437587. Archived from the original on 19 August 2020. Retrieved 2 December 2019.
	Rincon, Paul (21 August 2018). "Water ice 'detected on Moon's surface'". BBC News. Archived from the original on 21 August 2018. Retrieved 21 August 2018.
	David, Leonard. "Beyond the Shadow of a Doubt, Water Ice Exists on the Moon". Scientific American. Archived from the original on 21 August 2018. Retrieved 21 August 2018.
	"Water Ice Confirmed on the Surface of the Moon for the 1st Time!". Space.com. Archived from the original on 21 August 2018. Retrieved 21 August 2018.
	Honniball, C.I.; et al. (26 October 2020). "Molecular water detected on the sunlit Moon by SOFIA". Nature Astronomy. 5 (2): 121–127. Bibcode:2020NatAs.tmp..222H. doi:10.1038/s41550-020-01222-x. Archived from the original on 27 October 2020. Retrieved 26 October 2020.
	Hayne, P.O.; et al. (26 October 2020). "Micro cold traps on the Moon". Nature Astronomy. 5 (2): 169–175. arXiv:2005.05369. Bibcode:2020NatAs.tmp..221H. doi:10.1038/s41550-020-1198-9. S2CID 218595642. Archived from the original on 27 October 2020. Retrieved 26 October 2020.
	Guarino, Ben; Achenbach, Joel (26 October 2020). "Pair of studies confirm there is water on the moon - New research confirms what scientists had theorized for years — the moon is wet". The Washington Post. Archived from the original on 26 October 2020. Retrieved 26 October 2020.
	Chang, Kenneth (26 October 2020). "There's Water and Ice on the Moon, and in More Places Than NASA Once Thought - Future astronauts seeking water on the moon may not need to go into the most treacherous craters in its polar regions to find it". The New York Times. Archived from the original on 26 October 2020. Retrieved 26 October 2020.
	Schuerger, Andrew C.; Moores, John E.; Smith, David J.; Reitz, Günther (June 2019). "A Lunar Microbial Survival Model for Predicting the Forward Contamination of the Moon". Astrobiology. 19 (6): 730–756. Bibcode:2019AsBio..19..730S. doi:10.1089/ast.2018.1952. PMID 30810338.
	"Moon Storms". NASA. 27 September 2013. Archived from the original on 12 September 2013. Retrieved 3 October 2013.
	Culler, Jessica (16 June 2015). "LADEE - Lunar Atmosphere Dust and Environment Explorer". Archived from the original on 8 April 2015.
	Globus, Ruth (1977). "Chapter 5, Appendix J: Impact Upon Lunar Atmosphere". In Richard D. Johnson & Charles Holbrow (ed.). Space Settlements: A Design Study. NASA. Archived from the original on 31 May 2010. Retrieved 17 March 2010.
	Crotts, Arlin P.S. (2008). "Lunar Outgassing, Transient Phenomena and The Return to The Moon, I: Existing Data" (PDF). The Astrophysical Journal. 687 (1): 692–705. arXiv:0706.3949. Bibcode:2008ApJ...687..692C. doi:10.1086/591634. S2CID 16821394. Archived from the original (PDF) on 20 February 2009. Retrieved 29 September 2009.
	Steigerwald, William (17 August 2015). "NASA's LADEE Spacecraft Finds Neon in Lunar Atmosphere". NASA. Archived from the original on 19 August 2015. Retrieved 18 August 2015.
	Stern, S.A. (1999). "The Lunar atmosphere: History, status, current problems, and context". Reviews of Geophysics. 37 (4): 453–491. Bibcode:1999RvGeo..37..453S. CiteSeerX 10.1.1.21.9994. doi:10.1029/1999RG900005.
	Lawson, S.; Feldman, W.; Lawrence, D.; Moore, K.; Elphic, R.; Belian, R. (2005). "Recent outgassing from the lunar surface: the Lunar Prospector alpha particle spectrometer". Journal of Geophysical Research. 110 (E9): 1029. Bibcode:2005JGRE..11009009L. doi:10.1029/2005JE002433.
	R. Sridharan; S.M. Ahmed; Tirtha Pratim Dasa; P. Sreelathaa; P. Pradeepkumara; Neha Naika; Gogulapati Supriya (2010). "'Direct' evidence for water (H2O) in the sunlit lunar ambience from CHACE on MIP of Chandrayaan I". Planetary and Space Science. 58 (6): 947–950. Bibcode:2010P&SS...58..947S. doi:10.1016/j.pss.2010.02.013.
	"NASA: The Moon Once Had an Atmosphere That Faded Away". Time. Archived from the original on 14 October 2017. Retrieved 14 October 2017.
	Drake, Nadia; 17, National Geographic PUBLISHED June (17 June 2015). "Lopsided Cloud of Dust Discovered Around the Moon". National Geographic News. Archived from the original on 19 June 2015. Retrieved 20 June 2015.
	Horányi, M.; Szalay, J.R.; Kempf, S.; Schmidt, J.; Grün, E.; Srama, R.; Sternovsky, Z. (18 June 2015). "A permanent, asymmetric dust cloud around the Moon". Nature. 522 (7556): 324–326. Bibcode:2015Natur.522..324H. doi:10.1038/nature14479. PMID 26085272. S2CID 4453018.
	Matt Williams (10 July 2017). "How Long is a Day on the Moon?". Retrieved 5 December 2020.
	Haigh, I. D.; Eliot, M.; Pattiaratchi, C. (2011). "Global influences of the 18.61 year nodal cycle and 8.85 year cycle of lunar perigee on high tidal levels" (PDF). J. Geophys. Res. 116 (C6): C06025. Bibcode:2011JGRC..116.6025H. doi:10.1029/2010JC006645. Archived (PDF) from the original on 12 December 2019. Retrieved 24 September 2019.
	V V Belet︠s︡kiĭ (2001). Essays on the Motion of Celestial Bodies. Birkhäuser. p. 183. ISBN 978-3-7643-5866-2. Archived from the original on 23 March 2018. Retrieved 22 August 2020.
	Rambaux, N.; Williams, J. G. (2011). "The Moon's physical librations and determination of their free modes". Celestial Mechanics and Dynamical Astronomy. 109 (1): 85–100. Bibcode:2011CeMDA.109...85R. doi:10.1007/s10569-010-9314-2. S2CID 45209988.
	Amos, Jonathan (16 December 2009). "'Coldest place' found on the Moon". BBC News. Archived from the original on 11 August 2017. Retrieved 20 March 2010.
	"Diviner News". UCLA. 17 September 2009. Archived from the original on 7 March 2010. Retrieved 17 March 2010.
	Rocheleau, Jake (21 May 2012). "Temperature on the Moon – Surface Temperature of the Moon – PlanetFacts.org". Archived from the original on 27 May 2015.
	"Space Topics: Pluto and Charon". The Planetary Society. Archived from the original on 18 February 2012. Retrieved 6 April 2010.
	Phil Plait. "Dark Side of the Moon". Bad Astronomy: Misconceptions. Archived from the original on 12 April 2010. Retrieved 15 February 2010.
	Alexander, M.E. (1973). "The Weak Friction Approximation and Tidal Evolution in Close Binary Systems". Astrophysics and Space Science. 23 (2): 459–508. Bibcode:1973Ap&SS..23..459A. doi:10.1007/BF00645172. S2CID 122918899.
	"Moon used to spin 'on different axis'". BBC News. BBC. 23 March 2016. Archived from the original on 23 March 2016. Retrieved 23 March 2016.
	Luciuk, Mike. "How Bright is the Moon?". Amateur Astronomers. Archived from the original on 12 March 2010. Retrieved 16 March 2010.
	Hershenson, Maurice (1989). The Moon illusion. Routledge. p. 5. ISBN 978-0-8058-0121-7.
	Spekkens, K. (18 October 2002). "Is the Moon seen as a crescent (and not a "boat") all over the world?". Curious About Astronomy. Archived from the original on 16 October 2015. Retrieved 28 September 2015.
	"Moonlight helps plankton escape predators during Arctic winters". New Scientist. 16 January 2016. Archived from the original on 30 January 2016.
	"Supermoon November 2016". Space.com. 13 November 2016. Archived from the original on 14 November 2016. Retrieved 14 November 2016.
	Tony Phillips (16 March 2011). "Super Full Moon". NASA. Archived from the original on 7 May 2012. Retrieved 19 March 2011.
	Richard K. De Atley (18 March 2011). "Full moon tonight is as close as it gets". The Press-Enterprise. Archived from the original on 22 March 2011. Retrieved 19 March 2011.
	"'Super moon' to reach closest point for almost 20 years". The Guardian. 19 March 2011. Archived from the original on 25 December 2013. Retrieved 19 March 2011.
	Georgia State University, Dept. of Physics (Astronomy). "Perceived Brightness". Brightnes and Night/Day Sensitivity. Georgia State University. Archived from the original on 21 February 2014. Retrieved 25 January 2014.
	Lutron. "Measured light vs. perceived light" (PDF). From IES Lighting Handbook 2000, 27-4. Lutron. Archived (PDF) from the original on 5 February 2013. Retrieved 25 January 2014.
	Walker, John (May 1997). "Inconstant Moon". Earth and Moon Viewer. Fourth paragraph of "How Bright the Moonlight": Fourmilab. Archived from the original on 14 December 2013. Retrieved 23 January 2014. "14% [...] due to the logarithmic response of the human eye."
	Taylor, G. J. (8 November 2006). "Recent Gas Escape from the Moon". Planetary Science Research Discoveries: 110. Bibcode:2006psrd.reptE.110T. Archived from the original on 4 March 2007. Retrieved 4 April 2007.
	Schultz, P. H.; Staid, M. I.; Pieters, C. M. (2006). "Lunar activity from recent gas release". Nature. 444 (7116): 184–186. Bibcode:2006Natur.444..184S. doi:10.1038/nature05303. PMID 17093445. S2CID 7679109.
	"22 Degree Halo: a ring of light 22 degrees from the sun or moon". Department of Atmospheric Sciences, University of Illinois at Urbana–Champaign. Archived from the original on 5 April 2010. Retrieved 13 April 2010.
	Phillips, Tony (12 March 2007). "Stereo Eclipse". Science@NASA. Archived from the original on 10 June 2008. Retrieved 17 March 2010.
	Espenak, F. (2000). "Solar Eclipses for Beginners". MrEclip. Archived from the original on 24 May 2015. Retrieved 17 March 2010.
	Lambeck, K. (1977). "Tidal Dissipation in the Oceans: Astronomical, Geophysical and Oceanographic Consequences". Philosophical Transactions of the Royal Society A. 287 (1347): 545–594. Bibcode:1977RSPTA.287..545L. doi:10.1098/rsta.1977.0159. S2CID 122853694.
	Walker, John (10 July 2004). "Moon near Perigee, Earth near Aphelion". Fourmilab. Archived from the original on 8 December 2013. Retrieved 25 December 2013.
	Thieman, J.; Keating, S. (2 May 2006). "Eclipse 99, Frequently Asked Questions". NASA. Archived from the original on 11 February 2007. Retrieved 12 April 2007.
	Espenak, F. "Saros Cycle". NASA. Archived from the original on 24 May 2012. Retrieved 17 March 2010.
	Guthrie, D.V. (1947). "The Square Degree as a Unit of Celestial Area". Popular Astronomy. Vol. 55. pp. 200–203. Bibcode:1947PA.....55..200G.
	"Total Lunar Occultations". Royal Astronomical Society of New Zealand. Archived from the original on 23 February 2010. Retrieved 17 March 2010.
	Le Provost, C.; Bennett, A.F.; Cartwright, D.E. (1995). "Ocean Tides for and from TOPEX/POSEIDON". Science. 267 (5198): 639–642. Bibcode:1995Sci...267..639L. doi:10.1126/science.267.5198.639. PMID 17745840. S2CID 13584636.
	Touma, Jihad; Wisdom, Jack (1994). "Evolution of the Earth-Moon system". The Astronomical Journal. 108 (5): 1943–1961. Bibcode:1994AJ....108.1943T. doi:10.1086/117209.
	Chapront, J.; Chapront-Touzé, M.; Francou, G. (2002). "A new determination of lunar orbital parameters, precession constant and tidal acceleration from LLR measurements". Astronomy and Astrophysics. 387 (2): 700–709. Bibcode:2002A&A...387..700C. doi:10.1051/0004-6361:20020420. S2CID 55131241.
	"Why the Moon is getting further away from Earth". BBC News. 1 February 2011. Archived from the original on 25 September 2015. Retrieved 18 September 2015.
	Williams, James G.; Boggs, Dale H. (2016). "Secular tidal changes in lunar orbit and Earth rotation". Celestial Mechanics and Dynamical Astronomy. 126 (1): 89–129. Bibcode:2016CeMDA.126...89W. doi:10.1007/s10569-016-9702-3. ISSN 1572-9478. S2CID 124256137.
	Ray, R. (15 May 2001). "Ocean Tides and the Earth's Rotation". IERS Special Bureau for Tides. Archived from the original on 27 March 2010. Retrieved 17 March 2010.
	Stephenson, F. R.; Morrison, L. V.; Hohenkerk, C. Y. (2016). "Measurement of the Earth's rotation: 720 BC to AD 2015". Proceedings of the Royal Society A: Mathematical, Physical and Engineering Sciences. 472 (2196): 20160404. Bibcode:2016RSPSA.47260404S. doi:10.1098/rspa.2016.0404. PMC 5247521. PMID 28119545.
	Morrison, L. V.; Stephenson, F. R.; Hohenkerk, C. Y.; Zawilski, M. (2021). "Addendum 2020 to 'Measurement of the Earth's rotation: 720 BC to AD 2015'". Proceedings of the Royal Society A: Mathematical, Physical and Engineering Sciences. 477 (2246): 20200776. Bibcode:2021RSPSA.47700776M. doi:10.1098/rspa.2020.0776. S2CID 231938488.
	Murray, C.D.; Dermott, Stanley F. (1999). Solar System Dynamics. Cambridge University Press. p. 184. ISBN 978-0-521-57295-8.
	Dickinson, Terence (1993). From the Big Bang to Planet X. Camden East, Ontario: Camden House. pp. 79–81. ISBN 978-0-921820-71-0.
	Latham, Gary; Ewing, Maurice; Dorman, James; Lammlein, David; Press, Frank; Toksőz, Naft; Sutton, George; Duennebier, Fred; Nakamura, Yosio (1972). "Moonquakes and lunar tectonism". Earth, Moon, and Planets. 4 (3–4): 373–382. Bibcode:1972Moon....4..373L. doi:10.1007/BF00562004. S2CID 120692155.
	Iain Todd (31 March 2018). "Is the Moon maintaining Earth's magnetism?". BBC Sky at Night Magazine. Archived from the original on 22 September 2020. Retrieved 16 November 2020.
	"Lunar maps". Archived from the original on 1 June 2019. Retrieved 18 September 2019.
	"Carved and Drawn Prehistoric Maps of the Cosmos". Space Today. 2006. Archived from the original on 5 March 2012. Retrieved 12 April 2007.
	Aaboe, A.; Britton, J.P.; Henderson, J.A.; Neugebauer, Otto; Sachs, A.J. (1991). "Saros Cycle Dates and Related Babylonian Astronomical Texts". Transactions of the American Philosophical Society. 81 (6): 1–75. doi:10.2307/1006543. JSTOR 1006543. "One comprises what we have called "Saros Cycle Texts", which give the months of eclipse possibilities arranged in consistent cycles of 223 months (or 18 years)."
	Sarma, K.V. (2008). "Astronomy in India". In Helaine Selin (ed.). Encyclopaedia of the History of Science, Technology, and Medicine in Non-Western Cultures. Encyclopaedia of the History of Science (2 ed.). Springer. pp. 317–321. Bibcode:2008ehst.book.....S. ISBN 978-1-4020-4559-2.
	Needham, Joseph (1986). Science and Civilization in China, Volume III: Mathematics and the Sciences of the Heavens and Earth. Taipei: Caves Books. ISBN 978-0-521-05801-8. Archived from the original on 22 June 2019. Retrieved 22 August 2020.
	O'Connor, J.J.; Robertson, E.F. (February 1999). "Anaxagoras of Clazomenae". University of St Andrews. Archived from the original on 12 January 2012. Retrieved 12 April 2007.
	Robertson, E.F. (November 2000). "Aryabhata the Elder". Scotland: School of Mathematics and Statistics, University of St Andrews. Archived from the original on 11 July 2015. Retrieved 15 April 2010.
	A.I. Sabra (2008). "Ibn Al-Haytham, Abū ʿAlī Al-Ḥasan Ibn Al-Ḥasan". Dictionary of Scientific Biography. Detroit: Charles Scribner's Sons. pp. 189–210, at 195.
	Lewis, C.S. (1964). The Discarded Image. Cambridge: Cambridge University Press. p. 108. ISBN 978-0-521-47735-2. Archived from the original on 17 June 2020. Retrieved 11 November 2019.
	van der Waerden, Bartel Leendert (1987). "The Heliocentric System in Greek, Persian and Hindu Astronomy". Annals of the New York Academy of Sciences. 500 (1): 1–569. Bibcode:1987NYASA.500....1A. doi:10.1111/j.1749-6632.1987.tb37193.x. PMID 3296915. S2CID 84491987.
	Evans, James (1998). The History and Practice of Ancient Astronomy. Oxford & New York: Oxford University Press. pp. 71, 386. ISBN 978-0-19-509539-5.
	"Discovering How Greeks Computed in 100 B.C." The New York Times. 31 July 2008. Archived from the original on 4 December 2013. Retrieved 9 March 2014.
	Van Helden, A. (1995). "The Moon". Galileo Project. Archived from the original on 23 June 2004. Retrieved 12 April 2007.
	Consolmagno, Guy J. (1996). "Astronomy, Science Fiction and Popular Culture: 1277 to 2001 (And beyond)". Leonardo. 29 (2): 127–132. doi:10.2307/1576348. JSTOR 1576348. S2CID 41861791.
	Hall, R. Cargill (1977). "Appendix A: Lunar Theory Before 1964". NASA History Series. Lunar Impact: A History of Project Ranger. Washington, DC: Scientific and Technical Information Office, NASA. Archived from the original on 10 April 2010. Retrieved 13 April 2010.
	Zak, Anatoly (2009). "Russia's unmanned missions toward the Moon". Archived from the original on 14 April 2010. Retrieved 20 April 2010.
	"Rocks and Soils from the Moon". NASA. Archived from the original on 27 May 2010. Retrieved 6 April 2010.
	"Soldiers, Spies and the Moon: Secret U.S. and Soviet Plans from the 1950s and 1960s". The National Security Archive. National Security Archive. Archived from the original on 19 December 2016. Retrieved 1 May 2017.
	Brumfield, Ben (25 July 2014). "U.S. reveals secret plans for '60s moon base". CNN. Archived from the original on 27 July 2014. Retrieved 26 July 2014.
	Teitel, Amy (11 November 2013). "LUNEX: Another way to the Moon". Popular Science. Archived from the original on 16 October 2015.
	Logsdon, John (2010). John F. Kennedy and the Race to the Moon. Palgrave Macmillan. ISBN 978-0-230-11010-6.
	Coren, M. (26 July 2004). "'Giant leap' opens world of possibility". CNN. Archived from the original on 20 January 2012. Retrieved 16 March 2010.
	"Record of Lunar Events, 24 July 1969". Apollo 11 30th anniversary. NASA. Archived from the original on 8 April 2010. Retrieved 13 April 2010.
	"Manned Space Chronology: Apollo_11". Spaceline.org. Archived from the original on 14 February 2008. Retrieved 6 February 2008.
	"Apollo Anniversary: Moon Landing "Inspired World"". National Geographic. Archived from the original on 9 February 2008. Retrieved 6 February 2008.
	Orloff, Richard W. (September 2004) [First published 2000]. "Extravehicular Activity". Apollo by the Numbers: A Statistical Reference. NASA History Division, Office of Policy and Plans. The NASA History Series. Washington, DC: NASA. ISBN 978-0-16-050631-4. LCCN 00061677. NASA SP-2000-4029. Archived from the original on 6 June 2013. Retrieved 1 August 2013.
	Launius, Roger D. (July 1999). "The Legacy of Project Apollo". NASA History Office. Archived from the original on 8 April 2010. Retrieved 13 April 2010.
	SP-287 What Made Apollo a Success? A series of eight articles reprinted by permission from the March 1970 issue of Astronautics & Aeronautics, a publication of the American Institute of Aeronautics and Astronautics. Washington, DC: Scientific and Technical Information Office, National Aeronautics and Space Administration. 1971.
	"NASA news release 77-47 page 242" (PDF) (Press release). 1 September 1977. Archived (PDF) from the original on 4 June 2011. Retrieved 16 March 2010.
	Appleton, James; Radley, Charles; Deans, John; Harvey, Simon; Burt, Paul; Haxell, Michael; Adams, Roy; Spooner N.; Brieske, Wayne (1977). "NASA Turns A Deaf Ear To The Moon". OASI Newsletters Archive. Archived from the original on 10 December 2007. Retrieved 29 August 2007.
	Dickey, J.; Bender, P. L.; Faller, J. E.; Newhall, X. X.; Ricklefs, R. L.; Ries, J. G.; Shelus, P. J.; Veillet, C.; Whipple, A. L. (1994). "Lunar laser ranging: a continuing legacy of the Apollo program". Science. 265 (5171): 482–490. Bibcode:1994Sci...265..482D. doi:10.1126/science.265.5171.482. PMID 17781305. S2CID 10157934. Archived from the original on 19 August 2020. Retrieved 2 December 2019.
	"Hiten-Hagomoro". NASA. Archived from the original on 14 June 2011. Retrieved 29 March 2010.
	"Clementine information". NASA. 1994. Archived from the original on 25 September 2010. Retrieved 29 March 2010.
	"Lunar Prospector: Neutron Spectrometer". NASA. 2001. Archived from the original on 27 May 2010. Retrieved 29 March 2010.
	"SMART-1 factsheet". European Space Agency. 26 February 2007. Archived from the original on 23 March 2010. Retrieved 29 March 2010.
	"China's first lunar probe ends mission". Xinhua. 1 March 2009. Archived from the original on 4 March 2009. Retrieved 29 March 2010.
	David, Leonard (17 March 2015). "China Outlines New Rockets, Space Station and Moon Plans". Space.com. Archived from the original on 1 July 2016. Retrieved 29 June 2016.
	"KAGUYA Mission Profile". JAXA. Archived from the original on 28 March 2010. Retrieved 13 April 2010.
	"KAGUYA (SELENE) World's First Image Taking of the Moon by HDTV". Japan Aerospace Exploration Agency (JAXA) and Japan Broadcasting Corporation (NHK). 7 November 2007. Archived from the original on 16 March 2010. Retrieved 13 April 2010.
	"Mission Sequence". Indian Space Research Organisation. 17 November 2008. Archived from the original on 6 July 2010. Retrieved 13 April 2010.
	"Indian Space Research Organisation: Future Program". Indian Space Research Organisation. Archived from the original on 25 November 2010. Retrieved 13 April 2010.
	"India and Russia Sign an Agreement on Chandrayaan-2". Indian Space Research Organisation. 14 November 2007. Archived from the original on 17 December 2007. Retrieved 13 April 2010.
	"Lunar CRater Observation and Sensing Satellite (LCROSS): Strategy & Astronomer Observation Campaign". NASA. October 2009. Archived from the original on 1 January 2012. Retrieved 13 April 2010.
	"Giant moon crater revealed in spectacular up-close photos". NBC News. Space.com. 6 January 2012. Archived from the original on 18 March 2020. Retrieved 22 November 2019.
	Chang, Alicia (26 December 2011). "Twin probes to circle moon to study gravity field". Phys.org. Associated Press. Archived from the original on 22 July 2018. Retrieved 22 July 2018.
	Covault, C. (4 June 2006). "Russia Plans Ambitious Robotic Lunar Mission". Aviation Week. Archived from the original on 12 June 2006. Retrieved 12 April 2007.
	"About the Google Lunar X Prize". X-Prize Foundation. 2010. Archived from the original on 28 February 2010. Retrieved 24 March 2010.
	"President Bush Offers New Vision For NASA" (Press release). NASA. 14 December 2004. Archived from the original on 10 May 2007. Retrieved 12 April 2007.
	"Constellation". NASA. Archived from the original on 12 April 2010. Retrieved 13 April 2010.
	"NASA Unveils Global Exploration Strategy and Lunar Architecture" (Press release). NASA. 4 December 2006. Archived from the original on 23 August 2007. Retrieved 12 April 2007.
	Mann, Adam (July 2019). "NASA's Artemis Program". Space.com. Retrieved 19 April 2021.
	"India's Space Agency Proposes Manned Spaceflight Program". Space.com. 10 November 2006. Archived from the original on 11 April 2012. Retrieved 23 October 2008.
	"SpaceX to help Vodafone and Nokia install first 4G signal on the Moon | The Week UK". Archived from the original on 19 August 2020. Retrieved 28 February 2018.
	"NASA plans to send first woman on Moon by 2024". The Asian Age. 15 May 2019. Archived from the original on 14 April 2020. Retrieved 15 May 2019.
	Chang, Kenneth (24 January 2017). "For 5 Contest Finalists, a $20 Million Dash to the Moon". The New York Times. ISSN 0362-4331. Archived from the original on 15 July 2017. Retrieved 13 July 2017.
	Wall, Mike (16 August 2017), "Deadline for Google Lunar X Prize Moon Race Extended Through March 2018", space.com, archived from the original on 19 September 2017, retrieved 25 September 2017
	McCarthy, Ciara (3 August 2016). "US startup Moon Express approved to make 2017 lunar mission". The Guardian. ISSN 0261-3077. Archived from the original on 30 July 2017. Retrieved 13 July 2017.
	"An Important Update From Google Lunar XPRIZE". Google Lunar XPRIZE. 23 January 2018. Archived from the original on 24 January 2018. Retrieved 12 May 2018.
	"Moon Express Approved for Private Lunar Landing in 2017, a Space First". Space.com. Archived from the original on 12 July 2017. Retrieved 13 July 2017.
	Chang, Kenneth (29 November 2018). "NASA's Return to the Moon to Start With Private Companies' Spacecraft". The New York Times. Archived from the original on 1 December 2018. Retrieved 29 November 2018.
	Andrew Jones (23 September 2020). "China's Chang'e 3 lunar lander still going strong after 7 years on the moon". Archived from the original on 25 November 2020. Retrieved 16 November 2020.
	Takahashi, Yuki (September 1999). "Mission Design for Setting up an Optical Telescope on the Moon". California Institute of Technology. Archived from the original on 6 November 2015. Retrieved 27 March 2011.
	Chandler, David (15 February 2008). "MIT to lead development of new telescopes on moon". MIT News. Archived from the original on 4 March 2009. Retrieved 27 March 2011.
	Naeye, Robert (6 April 2008). "NASA Scientists Pioneer Method for Making Giant Lunar Telescopes". Goddard Space Flight Center. Archived from the original on 22 December 2010. Retrieved 27 March 2011.
	Bell, Trudy (9 October 2008). "Liquid Mirror Telescopes on the Moon". Science News. NASA. Archived from the original on 23 March 2011. Retrieved 27 March 2011.
	"Far Ultraviolet Camera/Spectrograph". Lpi.usra.edu. Archived from the original on 3 December 2013. Retrieved 3 October 2013.
	"Mission Report: Apollo 17 – The Most Productive Lunar Expedition" (PDF). NASA. Archived from the original (PDF) on 30 September 2006. Retrieved 10 February 2021.
	David, Leonard (21 October 2019). "Moon Dust Could Be a Problem for Future Lunar Explorers". Retrieved 26 November 2020.
	Zheng, William (15 January 2019). "Chinese lunar lander's cotton seeds spring to life on far side of the moon". South China Morning Post. Retrieved 26 November 2020.
	"Can any State claim a part of outer space as its own?". United Nations Office for Outer Space Affairs. Archived from the original on 21 April 2010. Retrieved 28 March 2010.
	"How many States have signed and ratified the five international treaties governing outer space?". United Nations Office for Outer Space Affairs. 1 January 2006. Archived from the original on 21 April 2010. Retrieved 28 March 2010.
	"Do the five international treaties regulate military activities in outer space?". United Nations Office for Outer Space Affairs. Archived from the original on 21 April 2010. Retrieved 28 March 2010.
	"Agreement Governing the Activities of States on the Moon and Other Celestial Bodies". United Nations Office for Outer Space Affairs. Archived from the original on 9 August 2010. Retrieved 28 March 2010.
	"The treaties control space-related activities of States. What about non-governmental entities active in outer space, like companies and even individuals?". United Nations Office for Outer Space Affairs. Archived from the original on 21 April 2010. Retrieved 28 March 2010.
	"Statement by the Board of Directors of the IISL On Claims to Property Rights Regarding The Moon and Other Celestial Bodies (2004)" (PDF). International Institute of Space Law. 2004. Archived from the original (PDF) on 22 December 2009. Retrieved 28 March 2010.
	"Further Statement by the Board of Directors of the IISL On Claims to Lunar Property Rights (2009)" (PDF). International Institute of Space Law. 22 March 2009. Archived from the original (PDF) on 22 December 2009. Retrieved 28 March 2010.
	Vazhapully, Kiran (22 July 2020). "Space Law at the Crossroads: Contextualizing the Artemis Accords and the Space Resources Executive Order". OpinioJuris. Retrieved 10 May 2021.
	"Administration Statement on Executive Order on Encouraging International Support for the Recovery and Use of Space Resources". SpaceRef.com. White House. Retrieved 17 June 2020.
	"Declaration of the Rights of the Moon". Australian Earth Laws Alliance. 11 February 2021. Retrieved 10 May 2021.
	Tepper, Eytan; Whitehead, Christopher (1 December 2018). "Moon, Inc.: The New Zealand Model of Granting Legal Personality to Natural Resources Applied to Space". New Space. 6 (4): 288–298. Bibcode:2018NewSp...6..288T. doi:10.1089/space.2018.0025. ISSN 2168-0256.
	Nemet-Nejat, Karen Rhea (1998), Daily Life in Ancient Mesopotamia, Daily Life, Greenwood, p. 203, ISBN 978-0-313-29497-6, archived from the original on 16 June 2020, retrieved 11 June 2019
	Black, Jeremy; Green, Anthony (1992). Gods, Demons and Symbols of Ancient Mesopotamia: An Illustrated Dictionary. The British Museum Press. p. 135. ISBN 978-0-7141-1705-8. Archived from the original on 19 August 2020. Retrieved 28 October 2017.
	Zschietzschmann, W. (2006). Hellas and Rome: The Classical World in Pictures. Whitefish, Montana: Kessinger Publishing. p. 23. ISBN 978-1-4286-5544-7.
	Cohen, Beth (2006). "Outline as a Special Technique in Black- and Red-figure Vase-painting". The Colors of Clay: Special Techniques in Athenian Vases. Los Angeles: Getty Publications. pp. 178–179. ISBN 978-0-89236-942-3. Archived from the original on 19 August 2020. Retrieved 28 April 2020.
	"Muhammad." Encyclopædia Britannica. 2007. Encyclopædia Britannica Online, p.13
	Ahead Of Chandrayaan 2 Landing, Poet-Diplomat Writes "Moon Anthem" Archived 20 September 2019 at the Wayback Machine NDTV, 6 Sept.2019
	Burton, David M. (2011). The History of Mathematics: An Introduction. Mcgraw-Hill. p. 3. ISBN 9780077419219.
	Brooks, A. S.; Smith, C. C. (1987). "Ishango revisited: new age determinations and cultural interpretations". The African Archaeological Review. 5: 65–78. doi:10.1007/BF01117083. JSTOR 25130482. S2CID 129091602.
	Duncan, David Ewing (1998). The Calendar. Fourth Estate Ltd. pp. 10–11. ISBN 978-1-85702-721-1.
	For etymology, see Barnhart, Robert K. (1995). The Barnhart Concise Dictionary of Etymology. Harper Collins. p. 487. ISBN 978-0-06-270084-1.. For the lunar calendar of the Germanic peoples, see Birley, A. R. (Trans.) (1999). Agricola and Germany. Oxford World's Classics. US: Oxford University Press. p. 108. ISBN 978-0-19-283300-6. Archived from the original on 17 June 2020. Retrieved 11 June 2019.
	Mallory, J.P.; Adams, D.Q. (2006). The Oxford Introduction to Proto-Indo-European and the Proto-Indo-European World. Oxford Linguistics. Oxford University Press. pp. 98, 128, 317. ISBN 978-0-19-928791-8.
	Harper, Douglas. "measure". Online Etymology Dictionary.
	Harper, Douglas. "menstrual". Online Etymology Dictionary.
	Smith, William George (1849). Dictionary of Greek and Roman Biography and Mythology: Oarses-Zygia. 3. J. Walton. p. 768. Archived from the original on 26 November 2020. Retrieved 29 March 2010.
	Estienne, Henri (1846). Thesaurus graecae linguae. 5. Didot. p. 1001. Archived from the original on 28 July 2020. Retrieved 29 March 2010.
	mensis. Charlton T. Lewis and Charles Short. A Latin Dictionary on Perseus Project.
	μείς in Liddell and Scott.
	Ilyas, Mohammad (March 1994). "Lunar Crescent Visibility Criterion and Islamic Calendar". Quarterly Journal of the Royal Astronomical Society. 35: 425. Bibcode:1994QJRAS..35..425L.
	Lilienfeld, Scott O.; Arkowitz, Hal (2009). "Lunacy and the Full Moon". Scientific American. Archived from the original on 16 October 2009. Retrieved 13 April 2010.
	Rotton, James; Kelly, I.W. (1985). "Much ado about the full moon: A meta-analysis of lunar-lunacy research". Psychological Bulletin. 97 (2): 286–306. doi:10.1037/0033-2909.97.2.286. PMID 3885282.
	Martens, R.; Kelly, I.W.; Saklofske, D.H. (1988). "Lunar Phase and Birthrate: A 50-year Critical Review". Psychological Reports. 63 (3): 923–934. doi:10.2466/pr0.1988.63.3.923. PMID 3070616. S2CID 34184527.
	Kelly, Ivan; Rotton, James; Culver, Roger (1986), "The Moon Was Full and Nothing Happened: A Review of Studies on the Moon and Human Behavior", Skeptical Inquirer, 10 (2): 129–143. Reprinted in The Hundredth Monkey - and other paradigms of the paranormal, edited by Kendrick Frazier, Prometheus Books. Revised and updated in The Outer Edge: Classic Investigations of the Paranormal, edited by Joe Nickell, Barry Karr, and Tom Genoni, 1996, CSICOP.
	
		Foster, Russell G.; Roenneberg, Till (2008). "Human Responses to the Geophysical Daily, Annual and Lunar Cycles". Current Biology. 18 (17): R784–R794. doi:10.1016/j.cub.2008.07.003. PMID 18786384. S2CID 15429616.
	
	Further reading
	
		Solar System portal Astronomy portal
	
		"Revisiting the Moon". The New York Times. Archived from the original on 8 September 2014. Retrieved 8 September 2014.
		"The Moon". Discovery 2008. BBC World Service. Retrieved 9 May 2021.
		Bussey, B.; Spudis, P.D. (2004). The Clementine Atlas of the Moon. Cambridge University Press. ISBN 978-0-521-81528-4.
		Cain, Fraser. "Where does the Moon Come From?". Universe Today. Retrieved 9 May 2021. (podcast and transcript)
		Jolliff, B. (2006). Wieczorek, M.; Shearer, C.; Neal, C. (eds.). New views of the Moon. Reviews in Mineralogy and Geochemistry. 60. Chantilly, Virginia: Mineralogy Society of America. p. 721. Bibcode:2006RvMG...60D...5J. doi:10.2138/rmg.2006.60.0. ISBN 978-0-939950-72-0. Archived from the original on 27 June 2007. Retrieved 12 April 2007.
		Jones, E. M. (2006). "Apollo Lunar Surface Journal". NASA. Retrieved 9 May 2021.
		"Exploring the Moon". Lunar and Planetary Institute. Retrieved 9 May 2021.
		Mackenzie, Dana (2003). The Big Splat, or How Our Moon Came to Be. Hoboken, NJ: John Wiley & Sons. ISBN 978-0-471-15057-2. Archived from the original on 17 June 2020. Retrieved 11 June 2019.
		Moore, P. (2001). On the Moon. Tucson, Arizona: Sterling Publishing Co. ISBN 978-0-304-35469-6.
		"Moon Articles". Planetary Science Research Discoveries. Hawai'i Institute of Geophysics and Planetology. Archived from the original on 17 November 2015. Retrieved 18 November 2006.
		Spudis, P.D. (1996). The Once and Future Moon. Smithsonian Institution Press. ISBN 978-1-56098-634-8. Archived from the original on 17 June 2020. Retrieved 11 June 2019.
		Taylor, S.R. (1992). Solar system evolution. Cambridge University Press. p. 307. ISBN 978-0-521-37212-1.
		Teague, K. (2006). "The Project Apollo Archive". Archived from the original on 4 April 2007. Retrieved 12 April 2007.
		Wilhelms, D.E. (1987). "Geologic History of the Moon". U.S. Geological Survey Professional Paper. Professional Paper. 1348. doi:10.3133/pp1348. Archived from the original on 23 February 2019. Retrieved 12 April 2007.
		Wilhelms, D.E. (1993). To a Rocky Moon: A Geologist's History of Lunar Exploration. Tucson: University of Arizona Press. ISBN 978-0-8165-1065-8. Archived from the original on 17 June 2020. Retrieved 10 March 2009.
	
	External links
	Moon
	at Wikipedia's sister projects
	
		Definitions from Wiktionary
		Media from Wikimedia Commons
		News from Wikinews
		Quotations from Wikiquote
		Texts from Wikisource
		Textbooks from Wikibooks
		Travel guides from Wikivoyage
		Resources from Wikiversity
	
		NASA images and videos about the Moon
		Albums of images and high-resolution overflight videos by Seán Doran, based on LROC data, on Flickr and YouTube
		Video (04:56) – The Moon in 4K (NASA, April 2018) on YouTube
		Video (04:47) – The Moon in 3D (NASA, July 2018) on YouTube
	
	Cartographic resources
	
		Unified Geologic Map of the Moon - United States Geological Survey
		Moon Trek – An integrated map browser of datasets and maps for the Moon
		The Moon on Google Maps, a 3-D rendition of the Moon akin to Google Earth
		"Consolidated Lunar Atlas". Lunar and Planetary Institute. Retrieved 26 February 2012.
		Gazetteer of Planetary Nomenclature (USGS) List of feature names.
		"Clementine Lunar Image Browser". U.S. Navy. 15 October 2003. Retrieved 12 April 2007.
		3D zoomable globes:
			"Google Moon". 2007. Retrieved 12 April 2007.
			"Moon". World Wind Central. NASA. 2007. Retrieved 12 April 2007.
		Aeschliman, R. "Lunar Maps". Planetary Cartography and Graphics. Retrieved 12 April 2007. Maps and panoramas at Apollo landing sites
		Japan Aerospace Exploration Agency (JAXA) Kaguya (Selene) images
		Lunar Earthside chart (4497 x 3150px)
		Large image of the Moon's north pole area Archived 23 August 2016 at the Wayback Machine
		Large image of Moon's south pole area (1000x1000px)
	
	Observation tools
	
		"NASA's SKYCAL – Sky Events Calendar". NASA. Archived from the original on 20 August 2007. Retrieved 27 August 2007.
		"Find moonrise, moonset and moonphase for a location". 2008. Retrieved 18 February 2008.
		"HMNAO's Moon Watch". 2005. Retrieved 24 May 2009. See when the next new crescent moon is visible for any location.
	
	General
	
		Lunar shelter (building a lunar base with 3D printing)
	
		vte
	
	The Moon
	
		vte
	
	Earth
	
		vte
	
	Natural satellites of the Solar System
	
		vte
	
	Solar System
	
		Stars portalSpaceflight portalOuter space portal
	
	Authority control Edit this at Wikidata
	Categories:
	
		Solar SystemMoonAstronomical objects known since antiquityMoonsPlanetary-mass satellitesPlanetary satellite systems
	
	Navigation menu
	
		Not logged in
		Talk
		Contributions
		Create account
		Log in
	
		Article
		Talk
	
		Read
		View source
		View history
	
	Search
	
		Main page
		Contents
		Current events
		Random article
		About Wikipedia
		Contact us
		Donate
	
	Contribute
	
		Help
		Learn to edit
		Community portal
		Recent changes
		Upload file
	
	Tools
	
		What links here
		Related changes
		Special pages
		Permanent link
		Page information
		Cite this page
		Wikidata item
	
	Print/export
	
		Download as PDF
		Printable version
	
	In other projects
	
		Wikimedia Commons
		Wikibooks
		Wikinews
		Wikiquote
		Wikisource
		Wikiversity
		Wikivoyage
	
	Languages
	
		Авар
		Башҡортса
		Чӑвашла
		Հայերեն
		日本語
		Нохчийн
		Русский
		Татарча/tatarça
		Удмурт
	
	Edit links
	
		This page was last edited on 22 September 2021, at 16:28 (UTC).
		Text is available under the Creative Commons Attribution-ShareAlike License; additional terms may apply. By using this site, you agree to the Terms of Use and Privacy Policy. Wikipedia® is a registered trademark of the Wikimedia Foundation, Inc., a non-profit organization.
	
		Privacy policy
		About Wikipedia
		Disclaimers
		Contact Wikipedia
		Mobile view
		Developers
		Statistics
		Cookie statement
	
		Wikimedia Foundation
		Powered by MediaWiki
	
	`

	us = `

	Wiki Loves Monuments: your chance to support Russian cultural heritage! Photograph a monument and win!
	Rings of Saturn
	From Wikipedia, the free encyclopedia
	Jump to navigation
	Jump to search
	For other uses, see Rings of Saturn (disambiguation).
	The full set of rings, imaged as Saturn eclipsed the Sun from the vantage of the Cassini orbiter, 1.2 million km distant, on 19 July 2013 (brightness is exaggerated). Earth appears as a dot at 4 o'clock, between the G and E rings.
	
	The rings of Saturn are the most extensive ring system of any planet in the Solar System. They consist of countless small particles, ranging in size from micrometers to meters,[1] that orbit around Saturn. The ring particles are made almost entirely of water ice, with a trace component of rocky material. There is still no consensus as to their mechanism of formation. Although theoretical models indicated that the rings were likely to have formed early in the Solar System's history,[2] new data from Cassini suggest they formed relatively late.[3]
	
	Although reflection from the rings increases Saturn's brightness, they are not visible from Earth with unaided vision. In 1610, the year after Galileo Galilei turned a telescope to the sky, he became the first person to observe Saturn's rings, though he could not see them well enough to discern their true nature. In 1655, Christiaan Huygens was the first person to describe them as a disk surrounding Saturn.[4] The concept that Saturn's rings are made up of a series of tiny ringlets can be traced to Pierre-Simon Laplace,[4] although true gaps are few – it is more correct to think of the rings as an annular disk with concentric local maxima and minima in density and brightness.[2] On the scale of the clumps within the rings there is much empty space.
	
	The rings have numerous gaps where particle density drops sharply: two opened by known moons embedded within them, and many others at locations of known destabilizing orbital resonances with the moons of Saturn. Other gaps remain unexplained. Stabilizing resonances, on the other hand, are responsible for the longevity of several rings, such as the Titan Ringlet and the G Ring.
	
	Well beyond the main rings is the Phoebe ring, which is presumed to originate from Phoebe and thus to share its retrograde orbital motion. It is aligned with the plane of Saturn's orbit. Saturn has an axial tilt of 27 degrees, so this ring is tilted at an angle of 27 degrees to the more visible rings orbiting above Saturn's equator.
	Voyager 2 view of Saturn casting a shadow across its rings. Four satellites, two of their shadows and ring spokes are visible.
	Contents
	
		1 History
			1.1 Early observations
			1.2 Huygens' ring theory and later developments
		2 Saturn's axial inclination
		3 Physical characteristics
		4 Formation and evolution of main rings
		5 Subdivisions and structures within the rings
			5.1 Physical parameters of the rings
				5.1.1 Major subdivisions
				5.1.2 C Ring structures
				5.1.3 Cassini Division structures
				5.1.4 A Ring structures
		6 D Ring
		7 C Ring
			7.1 Colombo Gap and Titan Ringlet
			7.2 Maxwell Gap and Ringlet
		8 B Ring
			8.1 Spokes
			8.2 Moonlet
		9 Cassini Division
			9.1 Huygens Gap
		10 A Ring
			10.1 Encke Gap
			10.2 Keeler Gap
			10.3 Propeller moonlets
		11 Roche Division
		12 F Ring
		13 Outer rings
			13.1 Janus/Epimetheus Ring
			13.2 G Ring
			13.3 Methone Ring Arc
			13.4 Anthe Ring Arc
			13.5 Pallene Ring
			13.6 E Ring
			13.7 Phoebe ring
		14 Possible ring system around Rhea
		15 Gallery
		16 See also
		17 Notes
		18 References
		19 External links
	
	History
	Early observations
	Detail of Galileo's drawing of Saturn in a letter to Belisario Vinta (1610).
	
	Galileo Galilei was the first to observe the rings of Saturn in 1610 using his telescope, but was unable to identify them as such. He wrote to the Duke of Tuscany that "The planet Saturn is not alone, but is composed of three, which almost touch one another and never move nor change with respect to one another. They are arranged in a line parallel to the zodiac, and the middle one (Saturn itself) is about three times the size of the lateral ones."[5] He also described the rings as Saturn's "ears". In 1612 the Earth passed through the plane of the rings and they became invisible. Mystified, Galileo remarked "I do not know what to say in a case so surprising, so unlooked for and so novel."[4] He mused, "Has Saturn swallowed his children?" — referring to the myth of the Titan Saturn devouring his offspring to forestall the prophecy of them overthrowing him.[5][6] He was further confused when the rings again became visible in 1613.[4]
	
	Early astronomers used anagrams as a form of commitment scheme to lay claim to new discoveries before their results were ready for publication. Galileo used smaismrmilmepoetaleumibunenugttauiras for Altissimum planetam tergeminum observavi ("I have observed the most distant planet to have a triple form") for discovering the rings of Saturn.[7]
	
	In 1657 Christopher Wren became Professor of Astronomy at Gresham College, London. He had been making observations of the planet Saturn from around 1652 with the aim of explaining its appearance. His hypothesis was written up in De corpore saturni, in which he came close to suggesting the planet had a ring. However, Wren was unsure whether the ring was independent of the planet, or physically attached to it. Before Wren's theory was published Christiaan Huygens presented his theory of the rings of Saturn. Immediately Wren recognised this as a better hypothesis than his own and De corpore saturni was never published. Robert Hooke was another early observer of the rings of Saturn, and noted the casting of shadows on the rings.[8]
	Huygens' ring theory and later developments
	Huygens' ring theory in Systema Saturnium (1659).
	
	Huygens began grinding lenses with his brother Constantijn in 1655 and was able to observe Saturn with greater detail using a 43× power refracting telescope that he designed himself. He was the first to suggest that Saturn was surrounded by a ring detached from the planet, and famously published the anagram: "aaaaaaacccccdeeeeeghiiiiiiillllmmnnnnnnnnnooooppqrrstttttuuuuu". Three years later, he revealed it to mean Annuto cingitur, tenui, plano, nusquam coherente, ad eclipticam inclinato ("[Saturn] is surrounded by a thin, flat, ring, nowhere touching, inclined to the ecliptic").[4][9] He published his ring theory in Systema Saturnium (1659) which also included his discovery of Saturn's moon, Titan, as well as the first clear outline of the dimensions of the Solar System.[10]
	
	In 1675, Giovanni Domenico Cassini determined that Saturn's ring was composed of multiple smaller rings with gaps between them; the largest of these gaps was later named the Cassini Division. This division is a 4,800-km-wide region between the A ring and B Ring.[11]
	
	In 1787, Pierre-Simon Laplace proved that a uniform solid ring would be unstable and suggested that the rings were composed of a large number of solid ringlets.[4][12]
	
	In 1859, James Clerk Maxwell demonstrated that a nonuniform solid ring, solid ringlets or a continuous fluid ring would also not be stable, indicating that the ring must be composed of numerous small particles, all independently orbiting Saturn.[12] Later, Sofia Kovalevskaya also found that Saturn's rings cannot be liquid ring-shaped bodies.[13] Spectroscopic studies of the rings carried out in 1895 by James Keeler of Allegheny Observatory and Aristarkh Belopolsky of Pulkovo Observatory showed Maxwell's analysis was correct.
	
	Four robotic spacecraft have observed Saturn's rings from the vicinity of the planet. Pioneer 11's closest approach to Saturn occurred in September 1979 at a distance of 20,900 km.[14] Pioneer 11 was responsible for the discovery of the F ring.[14] Voyager 1's closest approach occurred in November 1980 at a distance of 64,200 km.[15] A failed photopolarimeter prevented Voyager 1 from observing Saturn's rings at the planned resolution; nevertheless, images from the spacecraft provided unprecedented detail of the ring system and revealed the existence of the G ring.[16] Voyager 2's closest approach occurred in August 1981 at a distance of 41,000 km.[15] Voyager 2's working photopolarimeter allowed it to observe the ring system at higher resolution than Voyager 1, and to thereby discover many previously unseen ringlets.[17] Cassini spacecraft entered into orbit around Saturn in July 2004.[18] Cassini's images of the rings are the most detailed to-date, and are responsible for the discovery of yet more ringlets.[19]
	
	The rings are named alphabetically in the order they were discovered [20] (A and B in 1675 by Giovanni Domenico Cassini, C in 1850 by William Cranch Bond and his son George Phillips Bond, D in 1933 by Nikolai P. Barabachov and B. Semejkin, E in 1967 by Walter A. Feibelman, F in 1979 by Pioneer 11, and G in 1980 by Voyager 1). The main rings are, working outward from the planet, C, B and A, with the Cassini Division, the largest gap, separating Rings B and A. Several fainter rings were discovered more recently. The D Ring is exceedingly faint and closest to the planet. The narrow F Ring is just outside the A Ring. Beyond that are two far fainter rings named G and E. The rings show a tremendous amount of structure on all scales, some related to perturbations by Saturn's moons, but much unexplained.[20]
	Simulated appearance of Saturn as seen from Earth over the course of one Saturn year
	Saturn's axial inclination
	
	Saturn's axial tilt is 26.7°, meaning that widely varying views of the rings, of which the visible ones occupy its equatorial plane, are obtained from Earth at different times.[21] Earth makes passes through the ring plane every 13 to 15 years, about every half Saturn year, and there are about equal chances of either a single or three crossings occurring in each such occasion. The most recent ring plane crossings were on 22 May 1995, 10 August 1995, 11 February 1996 and 4 September 2009; upcoming events will occur on 23 March 2025, 15 October 2038, 1 April 2039 and 9 July 2039. Favorable ring plane crossing viewing opportunities (with Saturn not close to the Sun) only come during triple crossings.[22][23][24]
	
	Saturn's equinoxes, when the Sun passes through the ring plane, are not evenly spaced; on each orbit the Sun is south of the ring plane for 13.7 Earth years, then north of the plane for 15.7 years.[n 1] Dates for its northern hemisphere autumnal equinoxes include 19 November 1995 and 6 May 2025, with northern vernal equinoxes on 11 August 2009 and 23 January 2039.[26] During the period around an equinox the illumination of most of the rings is greatly reduced, making possible unique observations highlighting features that depart from the ring plane.[27]
	Physical characteristics
	Simulated image using color to present radio-occultation-derived particle size data. The attenuation of 0.94-, 3.6-, and 13-cm signals sent by Cassini through the rings to Earth shows abundance of particles of sizes similar to or larger than those wavelengths. Purple (B, inner A Ring) means few particles are < 5 cm (all signals similarly attenuated). Green and blue (C, outer A Ring) mean particles < 5 cm and < 1 cm, respectively, are common. White areas (B Ring) are too dense to transmit adequate signal. Other evidence shows rings A to C have a broad range of particle sizes, up to m across.
	The dark Cassini Division separates the wide inner B Ring and outer A ring in this image from the HST's ACS (March 22, 2004). The less prominent C Ring is just inside the B Ring.
	Cassini mosaic of Saturn's rings on August 12, 2009, a day after equinox. With the rings pointed at the Sun, illumination is by light reflected off Saturn, except on thicker or out-of-plane sections, like the F Ring.
	Cassini space probe view of the unilluminated side of Saturn's rings (May 9, 2007).
	
	The dense main rings extend from 7,000 km (4,300 mi) to 80,000 km (50,000 mi) away from Saturn's equator, whose radius is 60,300 km (37,500 mi) (see Major subdivisions). With an estimated local thickness of as little as 10 m[28] and as much as 1 km,[29] they are composed of 99.9% pure water ice with a smattering of impurities that may include tholins or silicates.[30] The main rings are primarily composed of particles ranging in size from 1 cm to 10 m.[31]
	
	Cassini directly measured the mass of the ring system via their gravitational effect during its final set of orbits that passed between the rings and the cloud tops, yielding a value of 1.54 (± 0.49) × 1019 kg, or 0.41 ± 0.13 Mimas masses.[3] This is as massive as about half the mass of the Earth's entire Antarctic ice shelf, spread across a surface area 80 times larger than that of Earth.[32] The estimate is close to the value of 0.40 Mimas masses derived from Cassini observations of density waves in the A, B and C rings.[3] It is a small fraction of the total mass of Saturn (about 0.25 ppb). Earlier Voyager observations of density waves in the A and B rings and an optical depth profile had yielded a mass of about 0.75 Mimas masses,[33] with later observations and computer modeling suggesting that was an underestimate.[34]
	
	Although the largest gaps in the rings, such as the Cassini Division and Encke Gap, can be seen from Earth, the Voyager spacecraft discovered that the rings have an intricate structure of thousands of thin gaps and ringlets. This structure is thought to arise, in several different ways, from the gravitational pull of Saturn's many moons. Some gaps are cleared out by the passage of tiny moonlets such as Pan,[35] many more of which may yet be discovered, and some ringlets seem to be maintained by the gravitational effects of small shepherd satellites (similar to Prometheus and Pandora's maintenance of the F ring). Other gaps arise from resonances between the orbital period of particles in the gap and that of a more massive moon further out; Mimas maintains the Cassini Division in this manner.[36] Still more structure in the rings consists of spiral waves raised by the inner moons' periodic gravitational perturbations at less disruptive resonances.[citation needed] Data from the Cassini space probe indicate that the rings of Saturn possess their own atmosphere, independent of that of the planet itself. The atmosphere is composed of molecular oxygen gas (O2) produced when ultraviolet light from the Sun interacts with water ice in the rings. Chemical reactions between water molecule fragments and further ultraviolet stimulation create and eject, among other things, O2. According to models of this atmosphere, H2 is also present. The O2 and H2 atmospheres are so sparse that if the entire atmosphere were somehow condensed onto the rings, it would be about one atom thick.[37] The rings also have a similarly sparse OH (hydroxide) atmosphere. Like the O2, this atmosphere is produced by the disintegration of water molecules, though in this case the disintegration is done by energetic ions that bombard water molecules ejected by Saturn's moon Enceladus. This atmosphere, despite being extremely sparse, was detected from Earth by the Hubble Space Telescope.[38] Saturn shows complex patterns in its brightness.[39] Most of the variability is due to the changing aspect of the rings,[40][41] and this goes through two cycles every orbit. However, superimposed on this is variability due to the eccentricity of the planet's orbit that causes the planet to display brighter oppositions in the northern hemisphere than it does in the southern.[42]
	
	In 1980, Voyager 1 made a fly-by of Saturn that showed the F ring to be composed of three narrow rings that appeared to be braided in a complex structure; it is now known that the outer two rings consist of knobs, kinks and lumps that give the illusion of braiding, with the less bright third ring lying inside them.[citation needed]
	
	New images of the rings taken around the 11 August 2009 equinox of Saturn by NASA's Cassini spacecraft have shown that the rings extend significantly out of the nominal ring plane in a few places. This displacement reaches as much as 4 km (2.5 mi) at the border of the Keeler Gap, due to the out-of-plane orbit of Daphnis, the moon that creates the gap.[43]
	Formation and evolution of main rings
	
	Estimates of the age of Saturn's rings vary widely, depending on the approach used. They have been considered to possibly be very old, dating to the formation of Saturn itself. However, data from Cassini suggest they are much younger, having most likely formed within the last 100 million years, and may thus be between 10 million and 100 million years old.[3][44] This recent origin scenario is based on a new, low mass estimate, modeling of the rings' dynamical evolution, and measurements of the flux of interplanetary dust, which feed into an estimate of the rate of ring darkening over time.[3] Since the rings are continually losing material, they would have been more massive in the past than at present.[3] The mass estimate alone is not very diagnostic, since high mass rings that formed early in the Solar System's history would have evolved by now to a mass close to that measured.[3] Based on current depletion rates, they may disappear in 300 million years.[45][46]
	
	There are two main theories regarding the origin of Saturn's inner rings. One theory, originally proposed by Édouard Roche in the 19th century, is that the rings were once a moon of Saturn (named Veritas, after a Roman goddess who hid in a well) whose orbit decayed until it came close enough to be ripped apart by tidal forces (see Roche limit).[47] A variation on this theory is that this moon disintegrated after being struck by a large comet or asteroid.[48] The second theory is that the rings were never part of a moon, but are instead left over from the original nebular material from which Saturn formed.[citation needed]
	A 2007 artist impression of the aggregates of icy particles that form the 'solid' portions of Saturn's rings. These elongated clumps are continually forming and dispersing. The largest particles are a few meters across.
	Saturn's rings
	and moons
	Tethys, Hyperion and Prometheus
	Tethys and Janus
	
	A more traditional version of the disrupted-moon theory is that the rings are composed of debris from a moon 400 to 600 km in diameter, slightly larger than Mimas. The last time there were collisions large enough to be likely to disrupt a moon that large was during the Late Heavy Bombardment some four billion years ago.[49]
	
	A more recent variant of this type of theory by R. M. Canup is that the rings could represent part of the remains of the icy mantle of a much larger, Titan-sized, differentiated moon that was stripped of its outer layer as it spiraled into the planet during the formative period when Saturn was still surrounded by a gaseous nebula.[50][51] This would explain the scarcity of rocky material within the rings. The rings would initially have been much more massive (≈1,000 times) and broader than at present; material in the outer portions of the rings would have coalesced into the moons of Saturn out to Tethys, also explaining the lack of rocky material in the composition of most of these moons.[51] Subsequent collisional or cryovolcanic evolution of Enceladus might then have caused selective loss of ice from this moon, raising its density to its current value of 1.61 g/cm3, compared to values of 1.15 for Mimas and 0.97 for Tethys.[51]
	
	The idea of massive early rings was subsequently extended to explain the formation of Saturn's moons out to Rhea.[52] If the initial massive rings contained chunks of rocky material (>100 km across) as well as ice, these silicate bodies would have accreted more ice and been expelled from the rings, due to gravitational interactions with the rings and tidal interaction with Saturn, into progressively wider orbits. Within the Roche limit, bodies of rocky material are dense enough to accrete additional material, whereas less-dense bodies of ice are not. Once outside the rings, the newly formed moons could have continued to evolve through random mergers. This process may explain the variation in silicate content of Saturn's moons out to Rhea, as well as the trend towards less silicate content closer to Saturn. Rhea would then be the oldest of the moons formed from the primordial rings, with moons closer to Saturn being progressively younger.[52]
	
	The brightness and purity of the water ice in Saturn's rings have also been cited as evidence that the rings are much younger than Saturn,[44] as the infall of meteoric dust would have led to a darkening of the rings. However, new research indicates that the B Ring may be massive enough to have diluted infalling material and thus avoided substantial darkening over the age of the Solar System. Ring material may be recycled as clumps form within the rings and are then disrupted by impacts. This would explain the apparent youth of some of the material within the rings.[53] Evidence suggesting a recent origin of the C ring has been gathered by researchers analyzing data from the Cassini Titan Radar Mapper, which focused on analyzing the proportion of rocky silicates within this ring. If much of this material was contributed by a recently disrupted centaur or moon, the age of this ring could be on the order of 100 million years or less. On the other hand, if the material came primarily from micrometeoroid influx, the age would be closer to a billion years.[54]
	
	The Cassini UVIS team, led by Larry Esposito, used stellar occultation to discover 13 objects, ranging from 27 m to 10 km across, within the F ring. They are translucent, suggesting they are temporary aggregates of ice boulders a few meters across. Esposito believes this to be the basic structure of the Saturnian rings, particles clumping together, then being blasted apart.[55]
	
	Research based on rates of infall into Saturn favors a younger ring system age of hundreds of millions of years. Ring material is continually spiraling down into Saturn; the faster this infall, the shorter the lifetime of the ring system. One mechanism involves gravity pulling electrically charged water ice grains down from the rings along planetary magnetic field lines, a process termed 'ring rain'. This flow rate was inferred to be 432–2870 kg/s using ground-based Keck telescope observations; as a consequence of this process alone, the rings will be gone in ~292+818
	−124 million years.[56] While traversing the gap between the rings and planet in September 2017, the Cassini spacecraft detected an equatorial flow of charge-neutral material from the rings to the planet of 4,800–44,000 kg/s.[57] Assuming this influx rate is stable, adding it to the continuous 'ring rain' process implies the rings may be gone in under 100 million years.[56][58]
	Subdivisions and structures within the rings
	
	The densest parts of the Saturnian ring system are the A and B Rings, which are separated by the Cassini Division (discovered in 1675 by Giovanni Domenico Cassini). Along with the C Ring, which was discovered in 1850 and is similar in character to the Cassini Division, these regions constitute the main rings. The main rings are denser and contain larger particles than the tenuous dusty rings. The latter include the D Ring, extending inward to Saturn's cloud tops, the G and E Rings and others beyond the main ring system. These diffuse rings are characterised as "dusty" because of the small size of their particles (often about a μm); their chemical composition is, like the main rings, almost entirely water ice. The narrow F Ring, just off the outer edge of the A Ring, is more difficult to categorize; parts of it are very dense, but it also contains a great deal of dust-size particles.
	Natural-color mosaic of Cassini narrow-angle camera images of the unilluminated side of Saturn's D, C, B, A and F rings (left to right) taken on May 9, 2007 (distances are to the planet's center).
	Physical parameters of the rings
	
	Notes:
	(1) Names as designated by the International Astronomical Union, unless otherwise noted. Broader separations between named rings are termed divisions, while narrower separations within named rings are called gaps.
	(2) Data mostly from the Gazetteer of Planetary Nomenclature, a NASA factsheet and several papers.[59][60][61]
	(3) distance is to centre of gaps, rings and ringlets that are narrower than 1,000 km
	(4) unofficial name
	The illuminated side of Saturn's rings with the major subdivisions labeled
	Major subdivisions
	Name(1) 	Distance from Saturn's
	center (km)(2) 	Width (km)(2) 	Named after
	D Ring 	66,900   –  74,510 	7,500 	 
	C Ring 	74,658   –   92,000 	17,500 	 
	B Ring 	92,000   –  117,580 	25,500 	 
	Cassini Division 	117,580   –   122,170 	4,700 	Giovanni Cassini
	A ring 	122,170   –   136,775 	14,600 	 
	Roche Division 	136,775   –   139,380 	2,600 	Édouard Roche
	F Ring 	140,180 (3) 	30   –  500 	 
	Janus/Epimetheus Ring(4) 	149,000   –  154,000 	5,000 	Janus and Epimetheus
	G Ring 	166,000   –  175,000 	9,000 	 
	Methone Ring Arc(4) 	194,230 	? 	Methone
	Anthe Ring Arc(4) 	197,665 	? 	Anthe
	Pallene Ring(4) 	211,000   –  213,500 	2,500 	Pallene
	E Ring 	180,000   –  480,000 	300,000 	 
	Phoebe Ring 	~4,000,000 – >13,000,000 		Phoebe  
	C Ring structures
	Name(1) 	Distance from Saturn's
	center (km)(2) 	Width (km)(2) 	Named after
	Colombo Gap 	77,870 (3) 	150 	Giuseppe "Bepi" Colombo
	Titan Ringlet 	77,870 (3) 	25 	Titan, moon of Saturn
	Maxwell Gap 	87,491 (3) 	270 	James Clerk Maxwell
	Maxwell Ringlet 	87,491 (3) 	64 	James Clerk Maxwell
	Bond Gap 	88,700 (3) 	30 	William Cranch Bond and George Phillips Bond
	1.470RS Ringlet 	88,716 (3) 	16 	its radius
	1.495RS Ringlet 	90,171 (3) 	62 	its radius
	Dawes Gap 	90,210 (3) 	20 	William Rutter Dawes
	Cassini Division structures
	
		Source:[62]
	
	Name(1) 	Distance from Saturn's
	center (km)(2) 	Width (km)(2) 	Named after
	Huygens Gap 	117,680 (3) 	285–400 	Christiaan Huygens
	Huygens Ringlet 	117,848 (3) 	~17 	Christiaan Huygens
	Herschel Gap 	118,234 (3) 	102 	William Herschel
	Russell Gap 	118,614 (3) 	33 	Henry Norris Russell
	Jeffreys Gap 	118,950 (3) 	38 	Harold Jeffreys
	Kuiper Gap 	119,405 (3) 	3 	Gerard Kuiper
	Laplace Gap 	119,967 (3) 	238 	Pierre-Simon Laplace
	Bessel Gap 	120,241 (3) 	10 	Friedrich Bessel
	Barnard Gap 	120,312 (3) 	13 	Edward Emerson Barnard
	A Ring structures
	Name(1) 	Distance from Saturn's
	center (km)(2) 	Width (km)(2) 	Named after
	Encke Gap 	133,589 (3) 	325 	Johann Encke
	Keeler Gap 	136,505 (3) 	35 	James Keeler
	Oblique (4 degree angle) Cassini images of Saturn's C, B, and A rings (left to right; the F ring is faintly visible in the full size upper image if viewed at sufficient brightness). Upper image: natural color mosaic of Cassini narrow-angle camera photos of the illuminated side of the rings taken on December 12, 2004. Lower image: simulated view constructed from a radio occultation observation conducted on May 3, 2005. Color in the lower image is used to represent information about ring particle sizes (see the caption of the article's second image for an explanation).
	D Ring
	A Cassini image of the faint D Ring, with the inner C Ring below
	
	The D Ring is the innermost ring, and is very faint. In 1980, Voyager 1 detected within this ring three ringlets designated D73, D72 and D68, with D68 being the discrete ringlet nearest to Saturn. Some 25 years later, Cassini images showed that D72 had become significantly broader and more diffuse, and had moved planetward by 200 km.[63]
	
	Present in the D Ring is a finescale structure with waves 30 km apart. First seen in the gap between the C Ring and D73,[63] the structure was found during Saturn's 2009 equinox to extend a radial distance of 19,000 km from the D Ring to the inner edge of the B Ring.[64][65] The waves are interpreted as a spiral pattern of vertical corrugations of 2 to 20 m amplitude;[66] the fact that the period of the waves is decreasing over time (from 60 km in 1995 to 30 km by 2006) allows a deduction that the pattern may have originated in late 1983 with the impact of a cloud of debris (with a mass of ≈1012 kg) from a disrupted comet that tilted the rings out of the equatorial plane.[63][64][67] A similar spiral pattern in Jupiter's main ring has been attributed to a perturbation caused by impact of material from Comet Shoemaker-Levy 9 in 1994.[64][68][69]
	C Ring
	View of the outer C Ring; the Maxwell Gap with the Maxwell Ringlet on its right side are above and right of center. The Bond Gap is above a broad light band towards the upper right; the Dawes Gap is within a dark band just below the upper right corner.
	
	The C Ring is a wide but faint ring located inward of the B Ring. It was discovered in 1850 by William and George Bond, though William R. Dawes and Johann Galle also saw it independently. William Lassell termed it the "Crepe Ring" because it seemed to be composed of darker material than the brighter A and B Rings.[70]
	
	Its vertical thickness is estimated at 5 m, its mass at around 1.1 × 1018 kg, and its optical depth varies from 0.05 to 0.12.[citation needed] That is, between 5 and 12 percent of light shining perpendicularly through the ring is blocked, so that when seen from above, the ring is close to transparent. The 30-km wavelength spiral corrugations first seen in the D Ring were observed during Saturn's equinox of 2009 to extend throughout the C Ring (see above).
	Colombo Gap and Titan Ringlet
	
	The Colombo Gap lies in the inner C Ring. Within the gap lies the bright but narrow Colombo Ringlet, centered at 77,883 km from Saturn's center, which is slightly elliptical rather than circular. This ringlet is also called the Titan Ringlet as it is governed by an orbital resonance with the moon Titan.[71] At this location within the rings, the length of a ring particle's apsidal precession is equal to the length of Titan's orbital motion, so that the outer end of this eccentric ringlet always points towards Titan.[71]
	Maxwell Gap and Ringlet
	
	The Maxwell Gap lies within the outer part of the C Ring. It also contains a dense non-circular ringlet, the Maxwell Ringlet. In many respects this ringlet is similar to the ε ring of Uranus. There are wave-like structures in the middle of both rings. While the wave in the ε ring is thought to be caused by Uranian moon Cordelia, no moon has been discovered in the Maxwell gap as of July 2008.[72]
	B Ring
	
	The B Ring is the largest, brightest, and most massive of the rings. Its thickness is estimated as 5 to 15 m and its optical depth varies from 0.4 to greater than 5,[73] meaning that >99% of the light passing through some parts of the B Ring is blocked. The B Ring contains a great deal of variation in its density and brightness, nearly all of it unexplained. These are concentric, appearing as narrow ringlets, though the B Ring does not contain any gaps.[citation needed]. In places, the outer edge of the B Ring contains vertical structures deviating up to 2.5 km from the main ring plane.
	
	A 2016 study of spiral density waves using stellar occultations indicated that the B Ring's surface density is in the range of 40 to 140 g/cm2, lower than previously believed, and that the ring's optical depth has little correlation with its mass density (a finding previously reported for the A and C rings).[73][74] The total mass of the B Ring was estimated to be somewhere in the range of 7 to 24×1018 kg. This compares to a mass for Mimas of 37.5×1018 kg.[73]
	High resolution (about 3 km per pixel) color view of the inner-central B Ring (98,600 to 105,500 km from Saturn's center). The structures shown (from 40 km wide ringlets at center to 300–500 km wide bands at right) remain sharply defined at scales below the resolution of the image.
	The B Ring's outer edge, viewed near equinox, where shadows are cast by vertical structures up to 2.5 km high, probably created by unseen embedded moonlets. The Cassini Division is at top.
	Spokes
	File:Saturn ring spokes PIA11144 secs15.5to23 20080926.ogvPlay media
	Dark spokes mark the B ring's sunlit side in low phase angle Cassini images. This is a low-bitrate video. Lo-res version of this video
	
	Until 1980, the structure of the rings of Saturn was explained as being caused exclusively by the action of gravitational forces. Then images from the Voyager spacecraft showed radial features in the B Ring, known as spokes,[75][76] which could not be explained in this manner, as their persistence and rotation around the rings was not consistent with gravitational orbital mechanics.[77] The spokes appear dark in backscattered light, and bright in forward-scattered light (see images in Gallery); the transition occurs at a phase angle near 60°. The leading theory regarding the spokes' composition is that they consist of microscopic dust particles suspended away from the main ring by electrostatic repulsion, as they rotate almost synchronously with the magnetosphere of Saturn. The precise mechanism generating the spokes is still unknown, although it has been suggested that the electrical disturbances might be caused by either lightning bolts in Saturn's atmosphere or micrometeoroid impacts on the rings.[77]
	
	The spokes were not observed again until some twenty-five years later, this time by the Cassini space probe. The spokes were not visible when Cassini arrived at Saturn in early 2004. Some scientists speculated that the spokes would not be visible again until 2007, based on models attempting to describe their formation. Nevertheless, the Cassini imaging team kept looking for spokes in images of the rings, and they were next seen in images taken on 5 September 2005.[78]
	
	The spokes appear to be a seasonal phenomenon, disappearing in the Saturnian midwinter and midsummer and reappearing as Saturn comes closer to equinox. Suggestions that the spokes may be a seasonal effect, varying with Saturn's 29.7-year orbit, were supported by their gradual reappearance in the later years of the Cassini mission.[79]
	Moonlet
	
	In 2009, during equinox, a moonlet embedded in the B ring was discovered from the shadow it cast. It is estimated to be 400 m (1,300 ft) in diameter.[80] The moonlet was given the provisional designation S/2009 S 1.
	Cassini Division
	The Cassini Division imaged from the Cassini spacecraft. The Huygens Gap lies at its right border; the Laplace Gap is towards the center. A number of other, narrower gaps are also present. The moon in the background is Mimas.
	
	The Cassini Division is a region 4,800 km (3,000 mi) in width between Saturn's A ring and B Ring. It was discovered in 1675 by Giovanni Cassini at the Paris Observatory using a refracting telescope that had a 2.5-inch objective lens with a 20-foot-long focal length and a 90x magnification.[81][82] From Earth it appears as a thin black gap in the rings. However, Voyager discovered that the gap is itself populated by ring material bearing much similarity to the C Ring.[72] The division may appear bright in views of the unlit side of the rings, since the relatively low density of material allows more light to be transmitted through the thickness of the rings (see second image in gallery).[citation needed]
	
	The inner edge of the Cassini Division is governed by a strong orbital resonance. Ring particles at this location orbit twice for every orbit of the moon Mimas.[83] The resonance causes Mimas' pulls on these ring particles to accumulate, destabilizing their orbits and leading to a sharp cutoff in ring density. Many of the other gaps between ringlets within the Cassini Division, however, are unexplained.[84][citation needed]
	Huygens Gap
	
	The Huygens Gap is located at the inner edge of the Cassini Division. It contains the dense, eccentric Huygens Ringlet in the middle. This ringlet exhibits irregular azimuthal variations of geometrical width and optical depth, which may be caused by the nearby 2:1 resonance with Mimas and the influence of the eccentric outer edge of the B-ring. There is an additional narrow ringlet just outside the Huygens Ringlet.[72]
	A Ring
	"A Ring" redirects here. For the letter, see Å.
	The central ringlet of the A Ring's Encke Gap coincides with Pan's orbit, implying its particles oscillate in horseshoe orbits.
	
	The A Ring is the outermost of the large, bright rings. Its inner boundary is the Cassini Division and its sharp outer boundary is close to the orbit of the small moon Atlas. The A Ring is interrupted at a location 22% of the ring width from its outer edge by the Encke Gap. A narrower gap 2% of the ring width from the outer edge is called the Keeler Gap.
	
	The thickness of the A Ring is estimated to be 10 to 30 m, its surface density from 35 to 40 g/cm2 and its total mass as 4 to 5×1018 kg[73] (just under the mass of Hyperion). Its optical depth varies from 0.4 to 0.9.[73]
	
	Similarly to the B Ring, the A Ring's outer edge is maintained by orbital resonances, albeit in this case a more complicated set. It is primarily acted on by the 7:6 resonance with Janus and Epimetheus, with other contributions from the 5:3 resonance with Mimas and various resonances with Prometheus and Pandora.[85][86] Other orbital resonances also excite many spiral density waves in the A Ring (and, to a lesser extent, other rings as well), which account for most of its structure. These waves are described by the same physics that describes the spiral arms of galaxies. Spiral bending waves, also present in the A Ring and also described by the same theory, are vertical corrugations in the ring rather than compression waves.[87]
	
	In April 2014, NASA scientists reported observing the possible formative stage of a new moon near the outer edge of the A Ring.[88][89]
	Encke Gap
	
	The Encke Gap is a 325-km-wide gap within the A ring, centered at a distance of 133,590 km from Saturn's center.[90] It is caused by the presence of the small moon Pan,[91] which orbits within it. Images from the Cassini probe have shown that there are at least three thin, knotted ringlets within the gap.[72] Spiral density waves visible on both sides of it are induced by resonances with nearby moons exterior to the rings, while Pan induces an additional set of spiraling wakes (see image in gallery).[72]
	
	Johann Encke himself did not observe this gap; it was named in honour of his ring observations. The gap itself was discovered by James Edward Keeler in 1888.[70] The second major gap in the A ring, discovered by Voyager, was named the Keeler Gap in his honor.[92]
	
	The Encke Gap is a gap because it is entirely within the A Ring. There was some ambiguity between the terms gap and division until the IAU clarified the definitions in 2008; before that, the separation was sometimes called the "Encke Division".[93]
	Keeler Gap
	Waves in the Keeler gap edges induced by the orbital motion of Daphnis (see also a stretched closeup view in the gallery).
	Near Saturn's equinox, Daphnis and its waves cast shadows on the A Ring.
	
	The Keeler Gap is a 42-km-wide gap in the A ring, approximately 250 km from the ring's outer edge. The small moon Daphnis, discovered 1 May 2005, orbits within it, keeping it clear.[94] The moon's passage induces waves in the edges of the gap (this is also influenced by its slight orbital eccentricity).[72] Because the orbit of Daphnis is slightly inclined to the ring plane, the waves have a component that is perpendicular to the ring plane, reaching a distance of 1500 m "above" the plane.[95][96]
	
	The Keeler gap was discovered by Voyager, and named in honor of the astronomer James Edward Keeler. Keeler had in turn discovered and named the Encke Gap in honor of Johann Encke.[70]
	Propeller moonlets
	Propeller moonlet Santos-Dumont from lit (top) and unlit sides of rings
	Location of the first four moonlets detected in the A ring.
	
	In 2006, four tiny "moonlets" were found in Cassini images of the A Ring.[97] The moonlets themselves are only about a hundred metres in diameter, too small to be seen directly; what Cassini sees are the "propeller"-shaped disturbances the moonlets create, which are several km across. It is estimated that the A Ring contains thousands of such objects. In 2007, the discovery of eight more moonlets revealed that they are largely confined to a 3,000 km belt, about 130,000 km from Saturn's center,[98] and by 2008 over 150 propeller moonlets had been detected.[99] One that has been tracked for several years has been nicknamed Bleriot.[100]
	Roche Division
	The Roche Division (passing through image center) between the A Ring and the narrow F Ring. Atlas can be seen within it. The Encke and Keeler gaps are also visible.
	
	The separation between the A ring and the F Ring has been named the Roche Division in honor of the French physicist Édouard Roche.[101] The Roche Division should not be confused with the Roche limit which is the distance at which a large object is so close to a planet (such as Saturn) that the planet's tidal forces will pull it apart.[102] Lying at the outer edge of the main ring system, the Roche Division is in fact close to Saturn's Roche limit, which is why the rings have been unable to accrete into a moon.[103]
	
	Like the Cassini Division, the Roche Division is not empty but contains a sheet of material.[citation needed] The character of this material is similar to the tenuous and dusty D, E, and G Rings.[citation needed] Two locations in the Roche Division have a higher concentration of dust than the rest of the region. These were discovered by the Cassini probe imaging team and were given temporary designations: R/2004 S 1, which lies along the orbit of the moon Atlas; and R/2004 S 2, centered at 138,900 km from Saturn's center, inward of the orbit of Prometheus.[104][105]
	F Ring
	File:PIA07712 - F ring animation.ogvPlay media
	The small moons Pandora (left) and Prometheus (right) orbit on either side of the F ring. Prometheus acts as a ring shepherd and is followed by dark channels that it has carved into the inner strands of the ring.
	
	The F Ring is the outermost discrete ring of Saturn and perhaps the most active ring in the Solar System, with features changing on a timescale of hours.[106] It is located 3,000 km beyond the outer edge of the A ring.[107] The ring was discovered in 1979 by the Pioneer 11 imaging team.[108] It is very thin, just a few hundred km in radial extent. While the traditional view has been that it is held together by two shepherd moons, Prometheus and Pandora, which orbit inside and outside it,[91] recent studies indicate that only Prometheus contributes to the confinement.[109][110] Numerical simulations suggest the ring was formed when Prometheus and Pandora collided with each other and were partially disrupted.[111]
	
	More recent closeup images from the Cassini probe show that the F Ring consists of one core ring and a spiral strand around it.[112] They also show that when Prometheus encounters the ring at its apoapsis, its gravitational attraction creates kinks and knots in the F Ring as the moon 'steals' material from it, leaving a dark channel in the inner part of the ring (see video link and additional F Ring images in gallery). Since Prometheus orbits Saturn more rapidly than the material in the F ring, each new channel is carved about 3.2 degrees in front of the previous one.[106]
	
	In 2008, further dynamism was detected, suggesting that small unseen moons orbiting within the F Ring are continually passing through its narrow core because of perturbations from Prometheus. One of the small moons was tentatively identified as S/2004 S 6.[106]
	A mosaic of 107 images showing 255° (about 70%) of the F Ring as it would appear if straightened out, showing the kinked primary strand and the spiral secondary strand. The radial width (top to bottom) is 1,500 km.
	Outer rings
	The outer rings seen back-illuminated by the Sun
	Janus/Epimetheus Ring
	
	A faint dust ring is present around the region occupied by the orbits of Janus and Epimetheus, as revealed by images taken in forward-scattered light by the Cassini spacecraft in 2006. The ring has a radial extent of about 5,000 km.[113] Its source is particles blasted off the moons' surfaces by meteoroid impacts, which then form a diffuse ring around their orbital paths.[114]
	G Ring
	
	The G Ring (see last image in gallery) is a very thin, faint ring about halfway between the F Ring and the beginning of the E Ring, with its inner edge about 15,000 km inside the orbit of Mimas. It contains a single distinctly brighter arc near its inner edge (similar to the arcs in the rings of Neptune) that extends about one-sixth of its circumference, centered on the half-km diameter moonlet Aegaeon, which is held in place by a 7:6 orbital resonance with Mimas.[115][116] The arc is believed to be composed of icy particles up to a few m in diameter, with the rest of the G Ring consisting of dust released from within the arc. The radial width of the arc is about 250 km, compared to a width of 9,000 km for the G Ring as a whole.[115] The arc is thought to contain matter equivalent to a small icy moonlet about a hundred m in diameter.[115] Dust released from Aegaeon and other source bodies within the arc by micrometeoroid impacts drifts outward from the arc because of interaction with Saturn's magnetosphere (whose plasma corotates with Saturn's magnetic field, which rotates much more rapidly than the orbital motion of the G Ring). These tiny particles are steadily eroded away by further impacts and dispersed by plasma drag. Over the course of thousands of years the ring gradually loses mass,[117] which is replenished by further impacts on Aegaeon.
	Methone Ring Arc
	
	A faint ring arc, first detected in September 2006, covering a longitudinal extent of about 10 degrees is associated with the moon Methone. The material in the arc is believed to represent dust ejected from Methone by micrometeoroid impacts. The confinement of the dust within the arc is attributable to a 14:15 resonance with Mimas (similar to the mechanism of confinement of the arc within the G ring).[118][119] Under the influence of the same resonance, Methone librates back and forth in its orbit with an amplitude of 5° of longitude.
	Anthe Ring Arc
	The Anthe Ring Arc – the bright spot is Anthe
	
	A faint ring arc, first detected in June 2007, covering a longitudinal extent of about 20 degrees is associated with the moon Anthe. The material in the arc is believed to represent dust knocked off Anthe by micrometeoroid impacts. The confinement of the dust within the arc is attributable to a 10:11 resonance with Mimas. Under the influence of the same resonance, Anthe drifts back and forth in its orbit over 14° of longitude.[118][119]
	Pallene Ring
	
	A faint dust ring shares Pallene's orbit, as revealed by images taken in forward-scattered light by the Cassini spacecraft in 2006.[113] The ring has a radial extent of about 2,500 km. Its source is particles blasted off Pallene's surface by meteoroid impacts, which then form a diffuse ring around its orbital path.[114][119]
	E Ring
	
	The E Ring is the second outermost ring and is extremely wide; it consists of many tiny (micron and sub-micron) particles of water ice with silicates, carbon dioxide and ammonia.[120] The E Ring is distributed between the orbits of Mimas and Titan.[121] Unlike the other rings, it is composed of microscopic particles rather than macroscopic ice chunks. In 2005, the source of the E Ring's material was determined to be cryovolcanic plumes[122][123] emanating from the "tiger stripes" of the south polar region of the moon Enceladus.[124] Unlike the main rings, the E Ring is more than 2,000 km thick and increases with its distance from Enceladus.[121] Tendril-like structures observed within the E Ring can be related to the emissions of the most active south polar jets of Enceladus.[125]
	
	Particles of the E Ring tend to accumulate on moons that orbit within it. The equator of the leading hemisphere of Tethys is tinted slightly blue due to infalling material.[126] The trojan moons Telesto, Calypso, Helene and Polydeuces are particularly affected as their orbits move up and down the ring plane. This results in their surfaces being coated with bright material that smooths out features.[127]
	The backlit E ring, with Enceladus silhouetted against it.
	The moon's south polar jets erupt brightly below it.
	Close-up of the south polar geysers of Enceladus, the source of the E Ring.
	Side view of Saturn system, showing Enceladus in relation to the E Ring
	E Ring tendrils from Enceladus geysers – comparison of images (a, c) with computer simulations
	Phoebe ring
	The Phoebe ring's huge extent dwarfs the main rings. Inset: 24 µm Spitzer image of part of the ring
	
	In October 2009, the discovery of a tenuous disk of material just interior to the orbit of Phoebe was reported. The disk was aligned edge-on to Earth at the time of discovery. This disk can be loosely described as another ring. Although very large (as seen from Earth, the apparent size of two full moons[128]), the ring is virtually invisible. It was discovered using NASA's infrared Spitzer Space Telescope,[129] and was seen over the entire range of the observations, which extended from 128 to 207 times the radius of Saturn,[130] with calculations indicating that it may extend outward up to 300 Saturn radii and inward to the orbit of Iapetus at 59 Saturn radii.[131] The ring was subsequently studied using the WISE, Herschel and Cassini spacecraft;[132] WISE observations show that it extends from at least between 50 and 100 to 270 Saturn radii (the inner edge is lost in the planet's glare).[133] Data obtained with WISE indicate the ring particles are small; those with radii greater than 10 cm comprise 10% or less of the cross-sectional area.[133]
	
	Phoebe orbits the planet at a distance ranging from 180 to 250 radii. The ring has a thickness of about 40 radii.[134] Because the ring's particles are presumed to have originated from impacts (micrometeoroid and larger) on Phoebe, they should share its retrograde orbit,[131] which is opposite to the orbital motion of the next inner moon, Iapetus. This ring lies in the plane of Saturn's orbit, or roughly the ecliptic, and thus is tilted 27 degrees from Saturn's equatorial plane and the other rings. Phoebe is inclined by 5° with respect to Saturn's orbit plane (often written as 175°, due to Phoebe's retrograde orbital motion), and its resulting vertical excursions above and below the ring plane agree closely with the ring's observed thickness of 40 Saturn radii.
	
	The existence of the ring was proposed in the 1970s by Steven Soter.[131] The discovery was made by Anne J. Verbiscer and Michael F. Skrutskie (of the University of Virginia) and Douglas P. Hamilton (of the University of Maryland, College Park).[130][135] The three had studied together at Cornell University as graduate students.[136]
	
	Ring material migrates inward due to reemission of solar radiation,[130] with a speed inversely proportional to particle size; a 3 cm particle would migrate from the vicinity of Phoebe to that of Iapetus over the age of the Solar System.[133] The material would thus strike the leading hemisphere of Iapetus. Infall of this material causes a slight darkening and reddening of the leading hemisphere of Iapetus (similar to what is seen on the Uranian moons Oberon and Titania) but does not directly create the dramatic two-tone coloration of that moon.[137] Rather, the infalling material initiates a positive feedback thermal self-segregation process of ice sublimation from warmer regions, followed by vapor condensation onto cooler regions. This leaves a dark residue of "lag" material covering most of the equatorial region of Iapetus's leading hemisphere, which contrasts with the bright ice deposits covering the polar regions and most of the trailing hemisphere.[138][139][140]
	Possible ring system around Rhea
	Main article: Rings of Rhea
	
	Saturn's second largest moon Rhea has been hypothesized to have a tenuous ring system of its own consisting of three narrow bands embedded in a disk of solid particles.[141][142] These putative rings have not been imaged, but their existence has been inferred from Cassini observations in November 2005 of a depletion of energetic electrons in Saturn's magnetosphere near Rhea. The Magnetospheric Imaging Instrument (MIMI) observed a gentle gradient punctuated by three sharp drops in plasma flow on each side of the moon in a nearly symmetric pattern. This could be explained if they were absorbed by solid material in the form of an equatorial disk containing denser rings or arcs, with particles perhaps several decimeters to approximately a meter in diameter. A more recent piece of evidence consistent with the presence of Rhean rings is a set of small ultraviolet-bright spots distributed in a line that extends three quarters of the way around the moon's circumference, within 2 degrees of the equator. The spots have been interpreted as the impact points of deorbiting ring material.[143] However, targeted observations by Cassini of the putative ring plane from several angles have turned up nothing, suggesting that another explanation for these enigmatic features is needed.[144]
	Gallery
	
		Saturn, behind the rings and draped with their shadows, as seen by Cassini from a distance of 725,000 km.
	
		Cassini image mosaic of the unlit side of the outer C Ring (bottom) and inner B Ring (top) near Saturn's equinox, showing multiple views of the shadow of Mimas. The shadow is attenuated by the denser B ring. The Maxwell Gap is below center.
	
		A spiral density wave in Saturn's inner B Ring which forms at a 2:1 orbital resonance with Janus. The wavelength decreases as the wave propagates away from the resonance, so the apparent foreshortening in the image is illusory.[n 2]
	
		Natural color view of the outer C Ring and B Ring.
	
		Dark B Ring spokes in a low-phase-angle Cassini image of the rings' unlit side. Left of center, two dark gaps (the larger being the Huygens Gap) and the bright (from this viewing geometry) ringlets to their left comprise the Cassini Division.
	
		Cassini image of the sun-lit side of the rings taken in 2009 at a phase angle of 144°, with bright B Ring spokes.
	
		Pan's motion through the A ring's Encke Gap induces edge waves and (non-self-propagating) spiraling wakes[145] ahead of and inward of it. The other more tightly wound bands are spiral density waves.
	
		Radially stretched (4x) view of the Keeler Gap edge waves induced by Daphnis.
	
		Prometheus (at right) and Pandora orbit just inside and outside the F Ring, but only Prometheus acts as a ring shepherd.
	
		Prometheus near apoapsis carving a dark channel in the F Ring (with older channels to the right). A movie of the process may be viewed at the Cassini Imaging Team website[146] or YouTube.[147]
	
		F ring dynamism, probably due to perturbing effects of small moonlets orbiting close to or through the ring's core.
	
		Saturn's shadow truncates the backlit G Ring and its bright inner arc. A video showing the arc's orbital motion may be viewed on YouTube[148] or the Cassini Imaging Team website.[149]
	
		Saturn and its A, B and C rings in visible and (inset) infrared light. In the false-color IR view, greater water ice content and larger grain size lead to blue-green color, while greater non-ice content and smaller grain size yield a reddish hue.
	
	See also
	
		Galileo Galilei – the first person to observe Saturn's rings, in 1610
		Christiaan Huygens – the first to propose that there was a ring surrounding Saturn, in 1655
		Giovanni Cassini – discovered the separation between the A and B rings (the Cassini Division), in 1675
		Édouard Roche – French astronomer who described how a satellite that comes within the Roche limit of Saturn could break up and form the rings
	
	Notes
	
	At 0.0565, Saturn's orbital eccentricity is the largest of the Solar System's giant planets, and over three times Earth's. Its aphelion is reached close to its northern hemisphere summer solstice.[25]
	
		Janus's orbital radius changes slightly each time it has a close encounter with its co-orbital moon Epimetheus. These encounters lead to periodic minor disruptions in the wave pattern.
	
	References
	
	Porco, Carolyn. "Questions around Saturn's rings". CICLOPS web site. Retrieved 2012-10-05.
	Tiscareno, M. S. (2012-07-04). "Planetary Rings". In Kalas, P.; French, L. (eds.). Planets, Stars and Stellar Systems. Springer. pp. 61–63. arXiv:1112.3305v2. doi:10.1007/978-94-007-5606-9_7. ISBN 978-94-007-5605-2. S2CID 118494597. Retrieved 2012-10-05.
	Iess, L.; Militzer, B.; Kaspi, Y.; Nicholson, P.; Durante, D.; Racioppa, P.; Anabtawi, A.; Galanti, E.; Hubbard, W.; Mariani, M. J.; Tortora, P.; Wahl, S.; Zannoni, M. (2019). "Measurement and implications of Saturn's gravity field and ring mass". Science. 364 (6445): eaat2965. Bibcode:2019Sci...364.2965I. doi:10.1126/science.aat2965. hdl:10150/633328. PMID 30655447. S2CID 58631177.
	Baalke, Ron. "Historical Background of Saturn's Rings". Saturn Ring Plane Crossings of 1995–1996. Jet Propulsion Laboratory. Archived from the original on 2009-03-21. Retrieved 2007-05-23.
	Whitehouse, David (2009). Renaissance Genius: Galileo Galilei and His Legacy to Modern Science. Sterling Publishing Company, Inc. p. 100. ISBN 978-1-4027-6977-1. OCLC 434563173.
	Deiss, B. M.; Nebel, V. (2016). "On a Pretended Observation of Saturn by Galileo". Journal for the History of Astronomy. 29 (3): 215–220. doi:10.1177/002182869802900301. S2CID 118636820.
	Miner, Ellis D.; et al. (2007). "The scientific significance of planetary ring systems". Planetary Ring Systems. Springer Praxis Books in Space Exploration. Praxis. pp. 1–16. doi:10.1007/978-0-387-73981-6_1. ISBN 978-0-387-34177-4.
	Alexander, A. F. O'D. (1962). The Planet Saturn. Quarterly Journal of the Royal Meteorological Society. 88. London: Faber and Faber Limited. pp. 108–109. Bibcode:1962QJRMS..88..366D. doi:10.1002/qj.49708837730. ISBN 978-0-486-23927-9.
	Campbell, John W., Jr. (April 1937). "Notes". Beyond the Life Line. Astounding Stories. pp. 81–85.
	"2004ESASP1278...11V Page 11". adsabs.harvard.edu.
	"Saturn's Cassini Division". StarChild. Retrieved 2007-07-06.
	"James Clerk Maxwell on the nature of Saturn's rings". JOC/EFR. March 2006. Retrieved 2007-07-08.
	"Kovalevsky, Sonya (or Kovalevskaya, Sofya Vasilyevna). Entry from Complete Dictionary of Scientific Biography". 2013.
	Dunford, Bill. "Pioneer 11 – In Depth". NASA web site. Archived from the original on 2015-12-08. Retrieved 2015-12-03.
	Angrum, Andrea. "Voyager – The Interstellar Mission". JPL/NASA web site. Retrieved 2015-12-03.
	Dunford, Bill. "Voyager 1 – In Depth". NASA web site. Archived from the original on 2015-10-03. Retrieved 2015-12-03.
	Dunford, Bill. "Voyager 2 – In Depth". NASA web site. Retrieved 2015-12-03.
	Dunford, Bill. "Cassini – Key Dates". NASA web site. Archived from the original on 2017-04-13. Retrieved 2015-12-03.
	Piazza, Enrico. "Cassini Solstice Mission: About Saturn & Its Moons". JPL/NASA web site. Retrieved 2015-12-03.
	"Solar System Exploration: Planets: Saturn: Rings". Solar System Exploration. Archived from the original on 2010-05-27.
	Williams, David R. (23 December 2016). "Saturn Fact Sheet". NASA. Archived from the original on 17 July 2017. Retrieved 12 October 2017.
	"Saturn Ring Plane Crossing 1995". pds.nasa.gov. NASA. 1997. Archived from the original on 2020-02-11. Retrieved 2020-02-11.
	"Hubble Views Saturn Ring-Plane Crossing". hubblesite.org. NASA. 5 June 1995. Archived from the original on 2020-02-11. Retrieved 2020-02-11.
	Lakdawalla, E. (2009-09-04). "Happy Saturn ring plane crossing day!". www.planetary.org/blogs. The Planetary Society. Retrieved 2020-02-11.
	Proctor, R.A. (1865). Saturn and Its System. London: Longman, Green, Longman, Roberts, & Green. p. 166. OCLC 613706938.
	Lakdawalla, E. (7 July 2016). "Oppositions, conjunctions, seasons, and ring plane crossings of the giant planets". planetary.org/blogs. The Planetary Society. Retrieved 17 February 2020.
	"PIA11667: The Rite of Spring". photojournal.jpl.nasa.gov. NASA/JPL. 21 September 2009. Retrieved 2020-02-17.
	Cornell University News Service (2005-11-10). "Researchers Find Gravitational Wakes In Saturn's Rings". ScienceDaily. Retrieved 2008-12-24.
	"Saturn: Rings". NASA. Archived from the original on 2010-05-27.
	Nicholson, P.D.; et al. (2008). "A close look at Saturn's rings with Cassini VIMS". Icarus. 193 (1): 182–212. Bibcode:2008Icar..193..182N. doi:10.1016/j.icarus.2007.08.036.
	Zebker, H.A.; et al. (1985). "Saturn's rings – Particle size distributions for thin layer model". Icarus. 64 (3): 531–548. Bibcode:1985Icar...64..531Z. doi:10.1016/0019-1035(85)90074-0.
	Koren, M. (2019-01-17). "The Massive Mystery of Saturn's Rings". The Atlantic. Retrieved 2019-01-21.
	Esposito, L. W.; O'Callaghan, M.; West, R. A. (1983). "The structure of Saturn's rings: Implications from the Voyager stellar occultation". Icarus. 56 (3): 439–452. Bibcode:1983Icar...56..439E. doi:10.1016/0019-1035(83)90165-3.
	Stewart, Glen R.; et al. (October 2007). "Evidence for a Primordial Origin of Saturn's Rings". Bulletin of the American Astronomical Society. American Astronomical Society, DPS meeting #39. 39: 420. Bibcode:2007DPS....39.0706S.
	Burns, J.A.; et al. (2001). "Dusty Rings and Circumplanetary Dust: Observations and Simple Physics" (PDF). In Grun, E.; Gustafson, B. A. S.; Dermott, S. T.; Fechtig H. (eds.). Interplanetary Dust. Berlin: Springer. pp. 641–725. Bibcode:2001indu.book..641B. ISBN 978-3-540-42067-5.
	Goldreich, Peter; et al. (1978). "The formation of the Cassini division in Saturn's rings". Icarus. 34 (2): 240–253. Bibcode:1978Icar...34..240G. doi:10.1016/0019-1035(78)90165-3.
	Rincon, Paul (2005-07-01). "Saturn rings have own atmosphere". British Broadcasting Corporation. Retrieved 2007-07-06.
	Johnson, R. E.; et al. (2006). "The Enceladus and OH Tori at Saturn" (PDF). The Astrophysical Journal. 644 (2): L137. Bibcode:2006ApJ...644L.137J. doi:10.1086/505750. S2CID 37698445. Archived from the original (PDF) on 2020-04-12.
	Schmude, Richard W Junior (2001). "Wideband photoelectric magnitude measurements of Saturn in 2000". Georgia Journal of Science. Retrieved 2007-10-14.
	Schmude, Richard, Jr. (2006-09-22). "Wideband photometric magnitude measurements of Saturn made during the 2005–06 Apparition". Georgia Journal of Science. ProQuest 230557408.
	Schmude, Richard W Jr (2003). "Saturn in 2002–03". Georgia Journal of Science. Retrieved 2007-10-14.
	Henshaw, C. (February 2003). "Variability in Saturn". Journal of the British Astronomical Association. British Astronomical Association. 113 (1). Retrieved 2017-12-20.
	"Surprising, Huge Peaks Discovered in Saturn's Rings". SPACE.com Staff. space.com. 2009-09-21. Retrieved 2009-09-26.
	Gohd, Chelsea (17 January 2019). "Saturn's rings are surprisingly young". Astronomy.com. Retrieved 2019-01-21.
	"NASA Research Reveals Saturn is Losing Its Rings at "Worst-Case-Scenario" Rate". Retrieved 2020-06-29.
	O'Donoghjue, James; et al. (April 2019). "Observations of the chemical and thermal response of 'ring rain' on Saturn's ionosphere". Icarus. 322: 251–206. Bibcode:2019Icar..322..251O. doi:10.1016/j.icarus.2018.10.027. hdl:2381/43180. Retrieved 2020-06-29.
	Baalke, Ron. "Historical Background of Saturn's Rings". 1849 Roche Proposes Tidal Break-up. Jet Propulsion Laboratory. Archived from the original on 2009-03-21. Retrieved 2008-09-13.
	"The Real Lord of the Rings". nasa.gov. 2002-02-12. Archived from the original on 2010-03-23.
	Kerr, Richard A (2008). "Saturn's Rings Look Ancient Again". Science. 319 (5859): 21. doi:10.1126/science.319.5859.21a. PMID 18174403. S2CID 30937575.
	Choi, C. Q. (2010-12-13). "Saturn's Rings Made by Giant "Lost" Moon, Study Hints". National Geographic. Retrieved 2012-11-05.
	Canup, R. M. (2010-12-12). "Origin of Saturn's rings and inner moons by mass removal from a lost Titan-sized satellite". Nature. 468 (7326): 943–6. Bibcode:2010Natur.468..943C. doi:10.1038/nature09661. PMID 21151108. S2CID 4326819.
	Charnoz, S.; et al. (December 2011). "Accretion of Saturn's mid-sized moons during the viscous spreading of young massive rings: Solving the paradox of silicate-poor rings versus silicate-rich moons". Icarus. 216 (2): 535–550. arXiv:1109.3360. Bibcode:2011Icar..216..535C. doi:10.1016/j.icarus.2011.09.017. S2CID 119222398.
	"Saturn's Rings May Be Old Timers". NASA/JPL and University of Colorado. 2007-12-12. Archived from the original on 2007-12-20. Retrieved 2008-01-24.
	Zhang, Z.; Hayes, A.G.; Janssen, M.A.; Nicholson, P.D.; Cuzzi, J.N.; de Pater, I.; Dunn, D.E.; Estrada, P.R.; Hedman, M.M. (2017). "Cassini microwave observations provide clues to the origin of Saturn's C ring". Icarus. 281: 297–321. Bibcode:2017Icar..281..297Z. doi:10.1016/j.icarus.2016.07.020.
	Esposito, L.W.; et al. (January 2012). "A predator–prey model for moon-triggered clumping in Saturn's rings". Icarus. 217 (1): 103–114. Bibcode:2012Icar..217..103E. doi:10.1016/j.icarus.2011.09.029.
	O’Donoghue, James; Moore, Luke; Connerney, Jack; Melin, Henrik; Stallard, Tom; Miller, Steve; Baines, Kevin H. (November 2018). "Observations of the chemical and thermal response of 'ring rain' on Saturn's ionosphere" (PDF). Icarus. 322: 251–260. Bibcode:2019Icar..322..251O. doi:10.1016/j.icarus.2018.10.027. hdl:2381/43180.
	Waite, J. H.; Perryman, R. S.; Perry, M. E.; Miller, K. E.; Bell, J.; Cravens, T. E.; Glein, C. R.; Grimes, J.; Hedman, M.; Cuzzi, J.; Brockwell, T.; Teolis, B.; Moore, L.; Mitchell, D. G.; Persoon, A.; Kurth, W. S.; Wahlund, J.-E.; Morooka, M.; Hadid, L. Z.; Chocron, S.; Walker, J.; Nagy, A.; Yelle, R.; Ledvina, S.; Johnson, R.; Tseng, W.; Tucker, O. J.; Ip, W.-H. (5 October 2018). "Chemical interactions between Saturn's atmosphere and its rings". Science. 362 (6410): eaat2382. Bibcode:2018Sci...362.2382W. doi:10.1126/science.aat2382. PMID 30287634.
	"Saturn is Officially Losing its Rings and Shockingly at Much Faster Rate than Expected". Sci-Tech Universe. Retrieved 2018-12-28.
	Porco, C.; et al. (October 1984). "The Eccentric Saturnian Ringlets at 1.29RS and 1.45RS". Icarus. 60 (1): 1–16. Bibcode:1984Icar...60....1P. doi:10.1016/0019-1035(84)90134-9.
	Porco, C. C.; et al. (November 1987). "Eccentric features in Saturn's outer C ring". Icarus. 72 (2): 437–467. Bibcode:1987Icar...72..437P. doi:10.1016/0019-1035(87)90185-0.
	Flynn, B. C.; et al. (November 1989). "Regular Structure in the Inner Cassini Division of Saturn's Rings". Icarus. 82 (1): 180–199. Bibcode:1989Icar...82..180F. doi:10.1016/0019-1035(89)90030-4.
	Lakdawalla, E. (2009-02-09). "New names for gaps in the Cassini Division within Saturn's rings". Planetary Society blog. Planetary Society. Retrieved 2017-12-20.
	Hedman, Matthew M.; et al. (2007). "Saturn's dynamic D ring" (PDF). Icarus. 188 (1): 89–107. Bibcode:2007Icar..188...89H. doi:10.1016/j.icarus.2006.11.017.
	Mason, J.; et al. (2011-03-31). "Forensic sleuthing ties ring ripples to impacts". CICLOPS press release. Cassini Imaging Central Laboratory for Operations. Retrieved 2011-04-04.
	"Extensive spiral corrugations". PIA 11664 caption. NASA / Jet Propulsion Laboratory / Space Science Institute. 2011-03-31. Retrieved 2011-04-04.
	"Tilting Saturn's rings". PIA 12820 caption. NASA / Jet Propulsion Laboratory / Space Science Institute. 2011-03-31. Retrieved 2011-04-04.
	Hedman, M. M.; et al. (2011-03-31). "Saturn's curiously corrugated C Ring". Science. 332 (6030): 708–11. Bibcode:2011Sci...332..708H. CiteSeerX 10.1.1.651.5611. doi:10.1126/science.1202238. PMID 21454753. S2CID 11449779.
	"Subtle Ripples in Jupiter's Ring". PIA 13893 caption. NASA / Jet Propulsion Laboratory-Caltech / SETI. 2011-03-31. Retrieved 2011-04-04.
	Showalter, M. R.; et al. (2011-03-31). "The impact of comet Shoemaker-Levy 9 sends ripples through the rings of Jupiter" (PDF). Science. 332 (6030): 711–3. Bibcode:2011Sci...332..711S. doi:10.1126/science.1202241. PMID 21454755. S2CID 27371440. Archived from the original (PDF) on 2020-02-12.
	Harland, David M., Mission to Saturn: Cassini and the Huygens Probe, Chichester: Praxis Publishing, 2002.
	Porco, C.; et al. (October 1984). "The eccentric Saturnian ringlets at 1.29Rs and 1.45Rs". Icarus. 60 (1): 1–16. Bibcode:1984Icar...60....1P. doi:10.1016/0019-1035(84)90134-9.
	Porco, C.C.; et al. (2005). "Cassini Imaging Science: Initial Results on Saturn'sRings and Small Satellites" (PDF). Science. 307 (5713): 1226–1236. Bibcode:2005Sci...307.1226P. doi:10.1126/science.1108056. PMID 15731439. S2CID 1058405.
	Hedman, M.M.; Nicholson, P.D. (2016-01-22). "The B-ring's surface mass density from hidden density waves: Less than meets the eye?". Icarus. 279: 109–124. arXiv:1601.07955. Bibcode:2016Icar..279..109H. doi:10.1016/j.icarus.2016.01.007. S2CID 119199474.
	Dyches, Preston (2 February 2016). "Saturn's Rings: Less than Meets the Eye?". NASA. Retrieved 3 February 2016.
	Smith, B. A.; Soderblom, L.; Batson, R.; Bridges, P.; Inge, J.; Masursky, H.; Shoemaker, E.; Beebe, R.; Boyce, J.; Briggs, G.; Bunker, A.; Collins, S. A.; Hansen, C. J.; Johnson, T. V.; Mitchell, J. L.; Terrile, R. J.; Cook Af, A. F.; Cuzzi, J.; Pollack, J. B.; Danielson, G. E.; Ingersoll, A. P.; Davies, M. E.; Hunt, G. E.; Morrison, D.; Owen, T.; Sagan, C.; Veverka, J.; Strom, R.; Suomi, V. E. (1982). "A New Look at the Saturn System: The Voyager 2 Images". Science. 215 (4532): 504–537. Bibcode:1982Sci...215..504S. doi:10.1126/science.215.4532.504. PMID 17771273. S2CID 23835071.
	"The Alphabet Soup of Saturn's Rings". The Planetary Society. 2007. Archived from the original on 2010-12-13. Retrieved 2007-07-24.
	Hamilton, Calvin (2004). "Saturn's Magnificent Rings". Retrieved 2007-07-25.
	Malik, Tarig (2005-09-15). "Cassini Probe Spies Spokes in Saturn's Rings". Imaginova Corp. Retrieved 2007-07-06.
	Mitchell, C.J.; et al. (2006). "Saturn's Spokes: Lost and Found" (PDF). Science. 311 (5767): 1587–9. Bibcode:2006Sci...311.1587M. CiteSeerX 10.1.1.368.1168. doi:10.1126/science.1123783. PMID 16543455. S2CID 36767835.
	"Cassini Solstice Mission: A Small Find Near Equinox". Cassini Solstice Mission. Archived from the original on 2009-10-10. Retrieved 2009-11-16.
	Webb, Thomas William (1859). Celestial Objects for Common Telescopes. Longman, Green, Longman, and Roberts. p. 130.
	Archie Frederick Collins, The greatest eye in the world: astronomical telescopes and their stories, page 8
	"Lecture 41: Planetary Rings". ohio-state.edu.
	O'Connor, J. J.; Robertson, E. F. (2003). "Giovanni Cassini - Biography". Maths History. School of Mathematics and Statistics University of St. Andrews, Scotland.
	El Moutamid et al 2015.
	Spahn, Frank; Hoffmann, Holger; Seiß, Martin; Seiler, Michael; Grätz, Fabio M. (19 June 2019). "The radial density profile of Saturn's A ring". arXiv:1906.08036 [astro-ph.EP].
	"Two Kinds of Wave". NASA Solar System Exploration. Retrieved 2019-05-30.
	Platt, Jane; et al. (14 April 2014). "NASA Cassini Images May Reveal Birth of a Saturn Moon". NASA.
	Murray, C. D.; Cooper, N. J.; Williams, G. A.; Attree, N. O.; Boyer, J. S. (2014-03-28). "The discovery and dynamical evolution of an object at the outer edge of Saturn's a ring". Icarus. 236: 165–168. Bibcode:2014Icar..236..165M. doi:10.1016/j.icarus.2014.03.024.
	Williams, David R. "Saturnian Rings Fact Sheet". NASA. Retrieved 2008-07-22.
	Esposito, L. W. (2002). "Planetary rings". Reports on Progress in Physics. 65 (12): 1741–1783. Bibcode:2002RPPh...65.1741E. doi:10.1088/0034-4885/65/12/201.
	Osterbrock, D. E.; Cruikshank, D. P. (1983). "J.E. Keeler's discovery of a gap in the outer part of the a ring". Icarus. 53 (2): 165. Bibcode:1983Icar...53..165O. doi:10.1016/0019-1035(83)90139-2.
	Blue, J. (2008-02-06). "Encke Division Changed to Encke Gap". USGS Astrogeology Science Center. USGS. Retrieved 2010-09-02.
	Porco, C.C.; et al. (2007). "Saturn's Small Inner Satellites: Clues to Their Origins" (PDF). Science. 318 (5856): 1602–1607. Bibcode:2007Sci...318.1602P. doi:10.1126/science.1143977. PMID 18063794. S2CID 2253135.
	Mason, Joe (11 June 2009). "Saturn's Approach To Equinox Reveals Never-before-seen Vertical Structures In Planet's Rings". CICLOPS web site. Retrieved 2009-06-13.
	Weiss, J. W.; et al. (11 June 2009). "Ring Edge Waves and the Masses of Nearby Satellites". The Astronomical Journal. 138 (1): 272–286. Bibcode:2009AJ....138..272W. CiteSeerX 10.1.1.653.4033. doi:10.1088/0004-6256/138/1/272.
	Tiscareno, Matthew S.; et al. (2006). "100-m-diameter moonlets in Saturn's A ring from observations of 'propeller' structures". Nature. 440 (7084): 648–650. Bibcode:2006Natur.440..648T. doi:10.1038/nature04581. PMID 16572165. S2CID 9688977.
	Sremčević, Miodrag; et al. (2007). "A belt of moonlets in Saturn's A ring". Nature. 449 (7165): 1019–1021. Bibcode:2007Natur.449.1019S. doi:10.1038/nature06224. PMID 17960236. S2CID 4330204.
	Tiscareno, Matthew S.; et al. (2008). "The population of propellers in Saturn's A Ring". Astronomical Journal. 135 (3): 1083–1091. arXiv:0710.4547. Bibcode:2008AJ....135.1083T. doi:10.1088/0004-6256/135/3/1083. S2CID 28620198.
	Porco, C. (2013-02-25). "Bleriot Recaptured". CICLOPS web site. NASA/JPL-Caltech/Space Science Institute. Retrieved 2013-02-27.
	"Planetary Names: Ring and Ring Gap Nomenclature". usgs.gov.
	Weisstein, Eric W. (2007). "Eric Weisstein's World of Physics – Roche Limit". scienceworld.wolfram.com. Retrieved 2007-09-05.
	NASA. "What is the Roche limit?". NASA–JPL. Archived from the original on 1999-11-05. Retrieved 2007-09-05.
	"IAUC 8401: S/2004 S 3, S/2004 S 4,, R/2004 S 1; 2004eg, 2004eh,, 2004ei". www.cbat.eps.harvard.edu.
	"IAUC 8432: Sats, RINGS OF SATURN; 2004fc". www.cbat.eps.harvard.edu.
	Murray, C. D.; et al. (June 5, 2008). "The determination of the structure of Saturn's F ring by nearby moonlets" (PDF). Nature. 453 (7196): 739–744. Bibcode:2008Natur.453..739M. doi:10.1038/nature06999. PMID 18528389. S2CID 205213483.
	Karttunen, H.; et al. (2007). Fundamental Astronomy. Springer-Verlag Berlin Heidelberg. ISBN 978-3-540-34144-4. OCLC 804078150. Retrieved 2013-05-25.
	Gehrels, T.; Baker, L. R.; Beshore, E.; Blenman, C.; Burke, J. J.; Castillo, N. D.; Dacosta, B.; Degewij, J.; Doose, L. R.; Fountain, J. W.; Gotobed, J.; Kenknight, C. E.; Kingston, R.; McLaughlin, G.; McMillan, R.; Murphy, R.; Smith, P. H.; Stoll, C. P.; Strickland, R. N.; Tomasko, M. G.; Wijesinghe, M. P.; Coffeen, D. L.; Esposito, L. (1980). "Imaging Photopolarimeter on Pioneer Saturn". Science. 207 (4429): 434–439. Bibcode:1980Sci...207..434G. doi:10.1126/science.207.4429.434. PMID 17833555. S2CID 25033550.
	Lakdawalla, E. (2014-07-05). "On the masses and motions of mini-moons: Pandora's not a "shepherd," but Prometheus still is". Planetary Society. Retrieved 2015-04-17.
	Cuzzi, J. N.; Whizin, A. D.; Hogan, R. C.; Dobrovolskis, A. R.; Dones, L.; Showalter, M. R.; Colwell, J. E.; Scargle, J. D. (April 2014). "Saturn's F Ring core: Calm in the midst of chaos". Icarus. 232: 157–175. Bibcode:2014Icar..232..157C. doi:10.1016/j.icarus.2013.12.027. ISSN 0019-1035.
	Hyodo, R.; Ohtsuki, K. (2015-08-17). "Saturn's F ring and shepherd satellites a natural outcome of satellite system formation". Nature Geoscience. 8 (9): 686–689. Bibcode:2015NatGe...8..686H. doi:10.1038/ngeo2508.
	Charnoz, S.; et al. (2005). "Cassini Discovers a Kinematic Spiral Ring Around Saturn" (PDF). Science. 310 (5752): 1300–1304. Bibcode:2005Sci...310.1300C. doi:10.1126/science.1119387. PMID 16311328. S2CID 6502280.
	NASA Planetary Photojournal PIA08328: Moon-Made Rings
	"NASA Finds Saturn's Moons May Be Creating New Rings". Cassini Legacy 1997–2007. Jet Propulsion Lab. 2006-10-11. Archived from the original on 2006-10-16. Retrieved 2017-12-20.
	Hedman, M. M.; et al. (2007). "The Source of Saturn's G Ring" (PDF). Science. 317 (5838): 653–656. Bibcode:2007Sci...317..653H. doi:10.1126/science.1143964. PMID 17673659. S2CID 137345.
	"S/2008 S 1. (NASA Cassini Saturn Mission Images)". ciclops.org.
	Davison, Anna (2 August 2007). "Saturn ring created by remains of long-dead moon". NewScientist.com news service.
	Porco C. C., [1]; et al. (2008-09-05). "More Ring Arcs for Saturn". Cassini Imaging Central Laboratory for Operations web site. Retrieved 2008-09-05.
	Hedman, M. M.; et al. (2008-11-25). "Three tenuous rings/arcs for three tiny moons". Icarus. 199 (2): 378–386. Bibcode:2009Icar..199..378H. doi:10.1016/j.icarus.2008.11.001.
	Hillier, JK; et al. (June 2007). "The composition of Saturn's E Ring". Monthly Notices of the Royal Astronomical Society. 377 (4): 1588–1596. Bibcode:2007MNRAS.377.1588H. doi:10.1111/j.1365-2966.2007.11710.x.
	Hedman, M. M.; et al. (2012). "The three-dimensional structure of Saturn's E Ring". Icarus. 217 (1): 322–338. arXiv:1111.2568. Bibcode:2012Icar..217..322H. doi:10.1016/j.icarus.2011.11.006. S2CID 1432112.
	Spahn, F.; et al. (2006-03-10). "Cassini Dust Measurements at Enceladus and Implications for the Origin of the E Ring". Science. 311 (5766): 1416–8. Bibcode:2006Sci...311.1416S. CiteSeerX 10.1.1.466.6748. doi:10.1126/science.1121375. PMID 16527969. S2CID 33554377.
	Porco, C. C.; Helfenstein, P.; Thomas, P. C.; Ingersoll, A. P.; Wisdom, J.; West, R.; Neukum, G.; Denk, T.; Wagner, R. (10 March 2006). "Cassini Observes the Active South Pole of Enceladus" (PDF). Science. 311 (5766): 1393–1401. Bibcode:2006Sci...311.1393P. doi:10.1126/science.1123013. PMID 16527964. S2CID 6976648.
	"Icy Tendrils Reaching into Saturn Ring Traced to Their Source". NASA News. 14 April 2015. Retrieved 2015-04-15.
	Mitchell, C. J.; Porco, C. C.; Weiss, J. W. (2015-04-15). "Tracking the geysers of Enceladus into Saturn's E ring" (PDF). The Astronomical Journal. 149 (5): 156. Bibcode:2015AJ....149..156M. doi:10.1088/0004-6256/149/5/156. ISSN 1538-3881. S2CID 55091776. Archived from the original (PDF) on 2019-03-08.
	Schenk Hamilton et al. 2011, pp. 751–53.
	Mason 2010.
	"NASA Space Telescope Discovers Largest Ring Around Saturn". NASA. July 3, 2017. Retrieved 2017-11-06.
	"JPL". NASA Jet Propulsion Laboratory (JPL).
	Verbiscer, Anne; et al. (2009-10-07). "Saturn's largest ring". Nature. 461 (7267): 1098–100. Bibcode:2009Natur.461.1098V. doi:10.1038/nature08515. PMID 19812546. S2CID 4349726.
	Cowen, Rob (2009-10-06). "Largest known planetary ring discovered". Science News.
	Tamayo, D.; et al. (2014-01-23). "First observations of the Phoebe ring in optical light". Icarus. 233: 1–8. arXiv:1401.6166. Bibcode:2014Icar..233....1T. doi:10.1016/j.icarus.2014.01.021. S2CID 40032407.
	Hamilton, Douglas P.; Skrutskie, Michael F.; Verbiscer, Anne J.; Masci, Frank J. (2015-06-10). "Small particles dominate Saturn's Phoebe ring to surprisingly large distances". Nature. 522 (7555): 185–187. Bibcode:2015Natur.522..185H. doi:10.1038/nature14476. PMID 26062508. S2CID 4464735.
	"The King of Rings". NASA, Spitzer Space Telescope center. 2009-10-07. Retrieved 2009-10-07.
	Grayson, Michelle (2009-10-07). "Huge 'ghost' ring discovered around Saturn". Nature News. doi:10.1038/news.2009.979.
	Weil, Martin (Oct 25, 2009). "U-Va., U-Md. astronomers find another Saturn ring". The Washington Post. p. 4C. Retrieved 2012-09-02.
	Denk, T.; et al. (2009-12-10). "Iapetus: Unique Surface Properties and a Global Color Dichotomy from Cassini Imaging" (PDF). Science. 327 (5964): 435–9. Bibcode:2010Sci...327..435D. doi:10.1126/science.1177088. PMID 20007863. S2CID 165865. Archived from the original (PDF) on 2020-02-27.
	"Cassini Is on the Trail of a Runaway Mystery". NASA Mission News. NASA. 8 October 2007. Retrieved 2017-12-20.
	Mason, J.; et al. (2009-12-10). "Cassini Closes In On The Centuries-old Mystery Of Saturn's Moon Iapetus". CICLOPS website newsroom. Space Science Institute. Retrieved 2009-12-22.
	Spencer, J. R.; et al. (2009-12-10). "Formation of Iapetus' Extreme Albedo Dichotomy by Exogenically Triggered Thermal Ice Migration". Science. 327 (5964): 432–5. Bibcode:2010Sci...327..432S. CiteSeerX 10.1.1.651.4218. doi:10.1126/science.1177132. PMID 20007862. S2CID 20663944.
	Jones, Geraint H.; et al. (2008-03-07). "The Dust Halo of Saturn's Largest Icy Moon, Rhea" (PDF). Science. 319 (5868): 1380–1384. Bibcode:2008Sci...319.1380J. doi:10.1126/science.1151524. PMID 18323452. S2CID 206509814. Archived from the original (PDF) on 2018-03-08.
	Lakdawalla, E. (2008-03-06). "A Ringed Moon of Saturn? Cassini Discovers Possible Rings at Rhea". The Planetary Society web site. Planetary Society. Archived from the original on March 10, 2008. Retrieved 2008-03-09.
	Lakdawalla, E. (5 October 2009). "Another possible piece of evidence for a Rhea ring". The Planetary Society Blog. Planetary Society. Archived from the original on 2012-02-17. Retrieved 2009-10-06.
	Kerr, Richard A. (2010-06-25). "The Moon Rings That Never Were". ScienceNow. Archived from the original on 2010-07-01. Retrieved 2010-08-05.
	NASA.gov
	"Soft Collision (NASA Cassini Saturn Mission Images)". ciclops.org.
	Prometheus collision. YouTube. 18 November 2007.
	Saturn's G Ring. YouTube. 6 August 2007.
	
		"Rounding the Corner (NASA Cassini Saturn Mission Images)". ciclops.org.
	
	External links
		Wikimedia Commons has media related to Rings of Saturn.
		Wikimedia Commons has media related to Rings of Saturn.
	
		Planetary Rings Node: Saturn's Ring System
		Saturn's Rings by NASA's Solar System Exploration
		Rings of Saturn nomenclature from the USGS planetary nomenclature page
		Biggest Ring Around Saturn Just Got Supersized (retrieved 2017-12-20 from Space.com)
		Everything a Curious Mind Should Know About Planetary Ring Systems with Dr Mark Showalter (Waseem Akhtar podcast with Mark Showalter)
		High-resolution animation by Seán Doran of the backlit rings
		High-resolution animation by Kevin M. Gill of a flyover of the outer B Ring at equinox (it starts getting less uniform after the first minute); see Rings album for more
		High-resolution animation by Nick Stevens of Saturn and its rings from an equatorial and a polar orbit, and from a dive below the rings; see listing for more
	
		vte
	
	Saturn
	
		vte
	
	Planetary rings
	
		vte
	
	Moons of Saturn
	
		vte
	
	Galileo Galilei
	
		vte
	
	Christiaan Huygens
	
		Astronomy portalStars portalSpaceflight portalOuter space portalSolar System portal
	
	Authority control Edit this at Wikidata
	Categories:
	
		Rings of SaturnPlanetary ringsAstronomical objects discovered in 1610Discoveries by Galileo GalileiDiscoveries by Christiaan Huygens
	
	Navigation menu
	
		Not logged in
		Talk
		Contributions
		Create account
		Log in
	
		Article
		Talk
	
		Read
		Edit
		View history
	
	Search
	
		Main page
		Contents
		Current events
		Random article
		About Wikipedia
		Contact us
		Donate
	
	Contribute
	
		Help
		Learn to edit
		Community portal
		Recent changes
		Upload file
	
	Tools
	
		What links here
		Related changes
		Special pages
		Permanent link
		Page information
		Cite this page
		Wikidata item
	
	Print/export
	
		Download as PDF
		Printable version
	
	In other projects
	
		Wikimedia Commons
	
	Languages
	
		Azərbaycanca
		Հայերեն
		हिन्दी
		日本語
		Русский
		Slovenčina
		Српски / srpski
		Suomi
		中文
	
	Edit links
	
		This page was last edited on 23 September 2021, at 21:33 (UTC).
		Text is available under the Creative Commons Attribution-ShareAlike License; additional terms may apply. By using this site, you agree to the Terms of Use and Privacy Policy. Wikipedia® is a registered trademark of the Wikimedia Foundation, Inc., a non-profit organization.
	
		Privacy policy
		About Wikipedia
		Disclaimers
		Contact Wikipedia
		Mobile view
		Developers
		Statistics
		Cookie statement
	
		Wikimedia Foundation
		Powered by MediaWiki
	
	`
)

func OriginalFrequency(s string) FreqMap {
	m := FreqMap{}
	for _, r := range s {
		m[r]++
	}
	return m
}

func TestConcurrentFrequency(t *testing.T) {
	seq := OriginalFrequency(euro + dutch + us)
	con := ConcurrentFrequency([]string{euro, dutch, us})
	if !reflect.DeepEqual(con, seq) {
		t.Fatal("ConcurrentFrequency wrong result")
	}
}

func TestSequentialFrequency(t *testing.T) {
	oSeq := OriginalFrequency(euro + dutch + us)
	seq := Frequency(euro + dutch + us)
	if !reflect.DeepEqual(oSeq, seq) {
		t.Fatal("Frequency wrong result")
	}
}

func BenchmarkSequentialFrequency(b *testing.B) {
	if testing.Short() {
		b.Skip("skipping benchmark in short mode.")
	}
	for i := 0; i < b.N; i++ {
		Frequency(euro + dutch + us)
	}
}

func BenchmarkConcurrentFrequency(b *testing.B) {
	if testing.Short() {
		b.Skip("skipping benchmark in short mode.")
	}
	for i := 0; i < b.N; i++ {
		ConcurrentFrequency([]string{euro, dutch, us})
	}
}
