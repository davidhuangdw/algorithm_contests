package code.jam.`2020`.qua
// https://codingcompetitions.withgoogle.com/codejam/round/000000000019fd27/000000000020bdf9

fun assign(intervals: List<List<Int>>): String{
    val points = intervals.flatMapIndexed{ i, v ->
        listOf(Triple(v[0], 0, i), Triple(v[1], 1, i))
    }.toMutableList()
    points.sortWith(compareBy({it.first}, {-it.second}, {it.third}))

    val remain = ArrayDeque(listOf('C', 'J'))
    val who = MutableList(intervals.size, {'-'})
    for((x, side, i) in points){
        if(side == 1){
            remain.add(who[i])
        } else if(remain.isEmpty()){
            return "IMPOSSIBLE"
        }else{
            val w = remain.removeFirst()
            who[i] = w
        }
    }
    return who.joinToString("")
}

fun main(){
    val T = readLine()!!.toInt()
    (1..T).forEach {
        val N = readLine()!!.toInt()
        val intervals = (1..N).map {
            readLine()!!.trim().split(" ").map{it.toInt()}
        }
        println("Case #${it}: ${assign(intervals)}")
    }
}
