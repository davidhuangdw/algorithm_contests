package atcoder.a7

fun main() {
    fun inputInts() = readLine()!!.split(' ').map(String::toInt)

    inputInts()
    val A = inputInts().toHashSet().sorted()

    fun count(p: Int): Int{
        val s = hashSetOf<Int>() // bug: shouldn't use booleanArray due to size
        var cnt = 0
        for(a in A){
            if(!s.contains(a % p)){
                cnt ++
                if(cnt >= 2){
                    return 2
                }
                s.add(a % p)
            }
        }
        return cnt
    }

    if(count(2) == 1){
        println(1)
        return
    }
    val d = A[1] - A[0]
    var r = d
    for(p in 2 until d){
        if(p * p > r){
            break
        }
        if(r % p == 0 && count(p) <= 1){
            println(1)
            return
        }
        while(r % p == 0){
            r /= p
        }
    }
    println(if(r > 1 && count(r) == 1) 1 else 2)
}
