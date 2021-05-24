package google_code_jam.`2020`.`1a`


fun solve(patterns: List<List<String>>): String{
    val pre = patterns.map{it.first()}.maxByOrNull{ it.length }!!
    val suf = patterns.map{it.last()}.maxByOrNull{ it.length }!!
    for(p in patterns){
        if(p[0] != pre.take(p[0].length) || p.last() != suf.takeLast(p.last().length)) return "*"
    }
    return buildString{
        append(pre)
        for(p in patterns)
            for(sub in p.slice(1..(p.size-2)))
                append(sub)
        append(suf)
    }
}

fun main(){
    val T = readLine()!!.toInt()
    (1..T).forEach {
        val N = readLine()!!.toInt()
        val patterns = (0 until N).map{ readLine()!!.trim().split("*")}
        println("Case #${it}: ${solve(patterns)}")
    }
}
