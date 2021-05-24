package google_code_jam.`2020`.`1a`

fun solve(N: Int): String{
    val poss = mutableListOf<Pair<Int, Int>>()
    val rem = N - 30
    if(rem < 0){
        for(i in 1..N)
            poss.add(i to 1)
    }else{
        var lside = true
        var left = N
        for(i in 1..30){
            if(rem and (1 shl i-1) != 0){
                val range = if(lside) 1..i else i.downTo(1)
                for(j in range)
                    poss.add(i to j)
                left -= 1 shl i-1
                lside = !lside
            }else{
                poss.add(if(lside) i to 1 else i to i)
                left -= 1
            }
        }
        for(i in 31..(30+left))
            poss.add(if(lside) i to 1 else i to i)
    }
    return buildString{
        for((i, j) in poss)
            appendLine("$i $j")
    }
}

fun main(){
    val T = readLine()!!.toInt()
    (1..T).forEach {
        val N = readLine()!!.toInt();
        println("Case #${it}: ")
        print(solve(N))
    }
}
